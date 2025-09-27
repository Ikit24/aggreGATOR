package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/Ikit24/aggreGATOR/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("follow command requires one argument")
	}
	ctx := context.Background()

	user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return err
	}
	url := cmd.Args[0]
	
	feed, err := s.db.GetFeedByURL(ctx, url)
	if err != nil {
		return err
	}
	now := time.Now().UTC()
	id := uuid.New()

	res, err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
        ID:        id,
        CreatedAt: now,
        UpdatedAt: now,
        UserID:    user.ID,
        FeedID:    feed.ID,
    })
    if err != nil {
        return err
    }

	fmt.Printf("%s %s\n", res.FeedName, res.UserName)
	return nil
}
