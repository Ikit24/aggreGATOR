package main

import (
	"context"
	"fmt"
	"time"
	"strings"
	"database/sql"

	"github.com/Ikit24/aggreGATOR/internal/database"
	"github.com/google/uuid"
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
	for ; ; <- ticker.C {
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
		desc := sql.NullString{
			String: item.Description,
			Valid: item.Description !="",
		}

		pub := sql.NullTime{}
		if t, ok := parsePubDate(item.PubDate); ok {
			pub = sql.NullTime{Time: t, Valid: true}
		}

		params := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: desc,
			PublishedAt: pub,
			FeedID:      feed.ID,
		}

		if err := s.db.CreatePost(ctx, params); err	!= nil{
			fmt.Println("create post error:", err)
			continue
		}
	}
	return nil
}

func parsePubDate(s string) (time.Time, bool) {
	layouts := []string{
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822Z,
		time.RFC822,
		time.RFC3339,
	}
	s = strings.TrimSpace(s)
	if s == "" {
		return time.Time{}, false
	}
	for _, layout := range layouts {
		t, err := time.Parse(layout, s)
		if err == nil {
			return t, true
		}
	}
	return time.Time{}, false
}
