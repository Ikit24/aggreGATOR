package main

import (
	"errors"
	"fmt"
	"context"

	"github.com/Ikit24/aggreGATOR/internal/database"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}
	return f(s, cmd)
}

func (c *commands) unfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("unfollow command requires one argument")
	}
	ctx := context.Background()
	url := cmd.Args[0]

	feed, err := s.db.GetFeedByURL(ctx, url)
	if err != nil {
		return err
	}
	
	if err := s.db.DeleteFeedFollowByUserAndFeed(ctx, database.DeleteFeedFollowByUserAndFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}); err != nil {
		return err
	}

	fmt.Printf("Unfollowed: %s\n", feed.Name)
	return nil
}
