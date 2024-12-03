package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gaschneider/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.params) != 2 {
		return fmt.Errorf("add feed handler expects two parameters for feed name and feed url")
	}

	name := cmd.params[0]
	url := cmd.params[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: name, Url: url, UserID: user.ID})
	if err != nil {
		return fmt.Errorf("error creating feed: %v", err)
	}

	fmt.Printf("%v\n", feed)

	_, err = addFeedFollow(s, feed.Url, user)

	return err
}
