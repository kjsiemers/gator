package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	//was: feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	feed, err := fetchFeed(context.Background(), "https://techcrunch.com/feed/")
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}
	fmt.Printf("Feed: %+v\n", feed)
	return nil
}
