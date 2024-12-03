package main

import "fmt"

type command struct {
	name   string
	params []string
}

type commands struct {
	allComands map[string]func(*state, command) error
}

func NewCommands() commands {
	newCommands := commands{}
	newCommands.allComands = make(map[string]func(*state, command) error, 0)

	return newCommands
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.allComands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	funcToRun, exists := c.allComands[cmd.name]
	if !exists {
		return fmt.Errorf("no '%v' command found", cmd.name)
	}

	return funcToRun(s, cmd)
}
