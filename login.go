package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.params) != 1 {
		return fmt.Errorf("login handler expects single parameter")
	}

	name := cmd.params[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("user doesn't exist")
	}

	err = s.config.SetUser(name)
	if err == nil {
		fmt.Println("User successfully logged")
	}
	return err
}
