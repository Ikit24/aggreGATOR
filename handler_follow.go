package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
    "github.com/lib/pq"
	"github.com/Ikit24/aggreGATOR/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
    if len(cmd.Args) != 1 {
        return fmt.Errorf("follow command requires one argument")
    }
    ctx := context.Background()
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
        if pqerr, ok := err.(*pq.Error); ok && pqerr.Code == "23505" {
            fmt.Printf("%s %s\n", feed.Name, user.Name)
            return nil
        }
        return err
    }

	fmt.Printf("%s %s\n", res.FeedName, res.UserName)
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
    ctx := context.Background()

    rows, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
    if err != nil {
        return err
    }

    for _, r := range rows {
        fmt.Println(r.FeedName)
    }
    return nil
}
