package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, _ command) error {
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error deleting users")
	}

	fmt.Println("Users successfully deleted")
	return err
}
