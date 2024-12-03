package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/gaschneider/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.params) != 1 {
		return fmt.Errorf("register handler expects single parameter for name")
	}

	name := cmd.params[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err == nil {
		return fmt.Errorf("user already registered")
	}

	if err != sql.ErrNoRows {
		return fmt.Errorf("error: %v", err)
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: cmd.params[0]})
	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}

	err = s.config.SetUser(name)
	if err == nil {
		fmt.Printf("User successfully created and logged: %v\n", user)
	}
	return err
}
