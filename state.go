package main

import (
	"github.com/gaschneider/blog_aggregator/internal/config"
	"github.com/gaschneider/blog_aggregator/internal/database"
)

type state struct {
	db     *database.Queries
	config *config.Config
}

func NewState() (state, error) {
	cfg, err := config.Read()
	if err != nil {
		return state{}, err
	}

	newState := state{}
	newState.config = &cfg

	return newState, nil
}

func (s *state) SetDB(db *database.Queries) {
	s.db = db
}
