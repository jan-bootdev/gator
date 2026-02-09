package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerUsers(s *state, cmd command) error {

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, u := range users {
		if u.Name == s.cfg.CurrentUserName {
			fmt.Println("* " + u.Name + " (current)")
		} else {
			fmt.Println("* " + u.Name)
		}
	}

	return nil
}

func handlerLogin(s *state, cmd command) error {

	fmt.Println(" -> login")

	user, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	return nil
}

func handlerRegister(s *state, cmd command) error {

	fmt.Printf("-> register %s\n", cmd)

	if len(cmd.args) != 1 {
		return fmt.Errorf("register requires 1 argument")
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	})
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	return nil
}
