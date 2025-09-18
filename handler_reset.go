package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if err := s.db.Reset(context.Background()); err	!= nil {
		return fmt.Errorf("reset failed: %w", err)
	}
	fmt.Println("Database reset successfully!")
	return nil
}

