package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, _ command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("user doesn't exist")
	}

	fmt.Println("Users:")
	for _, u := range users {
		if u.Name == s.config.CurrentUserName {
			fmt.Printf("* %v (current)\n", u.Name)
		} else {
			fmt.Printf("* %v\n", u.Name)
		}
	}

	return err
}
