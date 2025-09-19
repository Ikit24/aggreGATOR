package main

import (
	"fmt"
	"net/http"
	"encoding/xml"
	"io"
)

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return fmt.Errorf("couldn't fetch website: %s", err)
	}
	
	req.Header.Set("User-Agent", "gator")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}

	var rssFEED &RSSFeed
	err = xml.Unmarshal(body, &feed)

	return *RSSFeed, error
}
