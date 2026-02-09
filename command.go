package main

import (
	"gator/internal/config"
	"gator/internal/database"
)

type commands struct {
	handlers map[string]func(*state, command) error
}

type command struct {
	name string
	args []string
}

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func newCommands() commands {
	c := make(map[string]func(*state, command) error)
	return commands{c}
}

func (c *commands) run(s *state, cmd command) error {
	if c.handlers[cmd.name] == nil {
		return nil
	}

	return c.handlers[cmd.name](s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) error {
	c.handlers[name] = f
	return nil
}
