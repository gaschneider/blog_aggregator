package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, _ command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	fmt.Println("Feeds:")
	fmt.Println("Name - URL - User name")
	for _, f := range feeds {
		user, err := s.db.GetUserById(context.Background(), f.UserID)
		if err != nil {
			return fmt.Errorf("error: %v", err)
		}

		fmt.Printf("%v - %v - %v \n", f.Name, f.Url, user.Name)
	}

	return err
}
