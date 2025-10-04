package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Ikit24/aggreGATOR/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("agg command expects a single argument: time_between_reqs argument")
	}
	timeBetweenReqs := cmd.Args[0]
	timeParseDur, err := time.ParseDuration(timeBetweenReqs)
	if err != nil {
		return err
	}
	fmt.Printf("Collecting feeds every %s\n", timeParseDur)

	ticker := time.NewTicker(timeParseDur)
	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			fmt.Println("Error scraping feeds:", err)
		}
	}
}

func scrapeFeeds(s *state) error {
	ctx := context.Background()
	var err error
	var feed database.Feed
	var rssFeed *RSSFeed
	feed, err = s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return err
	}

	_, err = s.db.MarkFeedFetched(ctx, feed.ID)
	if err != nil {
		return err
	}

	rssFeed, err = fetchFeed(ctx, feed.Url)
	if err != nil {
		return err
	}

	for _, item := range rssFeed.Channel.Item {
		fmt.Printf("Title: %s\n", item.Title)
	}
	return nil
}
