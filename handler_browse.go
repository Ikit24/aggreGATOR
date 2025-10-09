package main

import (
	"context"
	"fmt"
	"time"
	"strconv"

	"github.com/Ikit24/aggreGATOR/internal/database"
)

func handlerBrowse(s *state, cmd command) error {
	// 1) parse limit
	limit := 2
	if len(cmd.Args) == 1 {
		n, err := strconv.Atoi(cmd.Args[0])
		if err != nil || n <= 0 {
			return fmt.Errorf("invalid limit")
		}
		limit = n
	} else if len(cmd.Args) > 1 {
		return fmt.Errorf("usage: %s [limit]", cmd.Name)
	}

	// 2) load current user (reuse your existing helper)
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	// 3) query posts
	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return err
	}

	// 4) print
	for _, p := range posts {
		pub := "unknown"
		if p.PublishedAt.Valid {
			pub = p.PublishedAt.Time.Format(time.RFC3339)
		}
		fmt.Printf("%s\n%s\npublished: %s\n\n", p.Title, p.Url, pub)
	}
	return nil
}
