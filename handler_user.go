package main

import (
	"context"
	"fmt"
	"time"

    "github.com/Ikit24/aggreGATOR/internal/database"
    "github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("login expects a single argument, the username.")
	}
	username := cmd.Args[0]
	
	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("user does not exist: %s", username)
	}
	if err := s.cfg.SetUser(username); err != nil {
		return err
	}
	fmt.Println("user set to:", username)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("register expects a single argument, the username.")
	}
	username := cmd.Args[0]
	
	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}

	_, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("user already exists: %s", username)
	}
	
	if err := s.cfg.SetUser(username); err != nil {
		return err
	}

	fmt.Println("user created:", username)
	return nil
}

func handlerListUsers(s *state, cmd command) error {
	ctx := context.Background()
	users, err := s.db.GetUsers(ctx)
	if err != nil {
		return fmt.Errorf("failed to list users")
	}

	for _, name := range users {
		if name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", name)
		} else {
			fmt.Printf("* %s\n", name)
		}
	}
	return nil
}
