package main

import (
	"context"
	"fmt"

	"github.com/gaschneider/blog_aggregator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.params) != 1 {
		return fmt.Errorf("unfollow handler expects single parameter for url")
	}

	url := cmd.params[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{UserID: user.ID, FeedID: feed.ID})
	if err == nil {
		fmt.Printf("Feed %v unfollowed\n", url)
	}
	return err
}
