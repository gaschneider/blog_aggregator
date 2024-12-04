package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/gaschneider/blog_aggregator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func getParams() ([]string, error) {
	if len(os.Args) < 2 {
		return nil, fmt.Errorf("no command executed")
	}

	return os.Args[1:], nil
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %v", err)
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req.Header.Add("User-Agent", "gator")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("performing request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %v", err)
	}

	var result RSSFeed
	if err := xml.Unmarshal(bodyBytes, &result); err != nil {
		return nil, fmt.Errorf("unmarshaling XML: %v", err)
	}

	result.Channel.Title = html.UnescapeString(result.Channel.Title)
	result.Channel.Description = html.UnescapeString(result.Channel.Description)
	for _, item := range result.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
	}

	return &result, nil
}

func addFeedFollow(s *state, feedUrl string, user database.User) (database.CreateFeedFollowRow, error) {
	feed, err := s.db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return database.CreateFeedFollowRow{}, fmt.Errorf("feed doesn't exist")
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), UserID: user.ID, FeedID: feed.ID,
	})

	return feedFollow, err

}

func autoDetectLayout(dateStr string) (time.Time, error) {
	layouts := []string{
		"Mon, 02 Jan 2006 15:04:05 -0700", // "Tue, 03 Dec 2024 15:06:34 +0000"
		"02 Jan 2006 15:04:05",            // "03 Dec 2024 15:06:34"
		"2006-01-02 15:04:05",             // "2024-12-03 15:06:34"
		"2006-01-02",                      // "2024-12-03"
		"15:04:05",                        // "15:06:34"
		"2006-01-02T15:04:05Z",            // "2024-12-03T15:06:34Z"
		"2006-01-02 15:04:05 -0700",       // "2024-12-03 15:06:34 +0000"
	}

	for _, layout := range layouts {
		parsedTime, err := time.Parse(layout, dateStr)
		if err == nil {
			return parsedTime, nil
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse time string: %s", dateStr)
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	rssFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}

	_, err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return err
	}

	for _, item := range rssFeed.Channel.Item {
		publishedDate, err := autoDetectLayout(item.PubDate)
		if err != nil {
			fmt.Printf("error parsing date %v\n", item.PubDate)
			continue
		}

		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(),
			Title: item.Title, Url: item.Link, Description: item.Description, PublishedAt: publishedDate, FeedID: nextFeed.ID,
		})
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
				pattern := `"(posts_url_key)"`
				re := regexp.MustCompile(pattern)
				matches := re.FindStringSubmatch(pqErr.Message)
				if len(matches) > 1 {
					continue
				}
			}
			fmt.Printf("error: %v\n", err)
		}
	}

	return nil
}
