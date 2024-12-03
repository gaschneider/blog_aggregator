package main

import (
	"context"
	"fmt"

	"github.com/gaschneider/blog_aggregator/internal/database"
)

func handlerFollowing(s *state, _ command, user database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	fmt.Println("Following feeds:")
	for _, f := range feeds {
		fmt.Printf("%v \n", f.FeedName)
	}

	return nil
}
