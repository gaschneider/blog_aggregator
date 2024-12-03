package main

import (
	"fmt"

	"github.com/gaschneider/blog_aggregator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.params) != 1 {
		return fmt.Errorf("follow handler expects single parameter for url")
	}

	url := cmd.params[0]

	_, err := addFeedFollow(s, url, user)

	return err
}
