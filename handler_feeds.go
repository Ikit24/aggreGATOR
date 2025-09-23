package main

import (
	"fmt"
	"context"
)

func handlerFeeds(s *state, cmd command) error {
	ctx := context.Background()

	rows, err := s.db.ListFeedsWithUsers(ctx)
	if err != nil {
		return err
	}

	for _, r := range rows {
		fmt.Println(r.Name)
		fmt.Println(r.Url)
		fmt.Println(r.CreatorName)
		fmt.Println()
    }
    return nil
}
