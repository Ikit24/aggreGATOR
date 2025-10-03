package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	ctx := context.Background()
	feed, err := fetchFeed(ctx, "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Printf("%v+\n", feed)
	return nil
}

func scrapeFeeds(s *state) error {
	ctx := context.Background()
	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return err
	}

	_, err := s.db.MarkFeedFetched(ctx, feed.ID)
	if err != nil {
		return err
	}

	rssFeed, err := fetchFeed(ctx, feed.URL)
	if err != nil {
		return err
	}

	for _, item := range rssFeed.Channel.Item {
		fmt.Printf("Title: %s\n", item.Title)
	}
	return nil
}
