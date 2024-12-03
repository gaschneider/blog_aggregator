package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gaschneider/blog_aggregator/internal/database"
	"github.com/google/uuid"
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
		fmt.Printf("%v\n", item.Title)
	}

	return nil
}
