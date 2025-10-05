package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	ctx := context.Background()
	if err := s.db.DeleteFeedFollows(ctx); err != nil {
		return fmt.Errorf("reset failed: %w", err)
	}
	if err := s.db.DeleteFeeds(ctx); err != nil {
		return fmt.Errorf("reset failed: %w", err)
	}
	if err := s.db.DeleteUsers(ctx); err != nil {
		return fmt.Errorf("reset failed: %w", err)
	}
	fmt.Println("Database reset successfully!")
	return nil
}

