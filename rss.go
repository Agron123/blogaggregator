package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gator")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	rss := RSSFeed{}
	if err = xml.Unmarshal(bytes, &rss); err != nil {
		return nil, err
	}

	rss.Channel.Description = html.UnescapeString(rss.Channel.Description)
	rss.Channel.Title = html.UnescapeString(rss.Channel.Title)

	for index, _ := range rss.Channel.Item {
		rss.Channel.Item[index].Title = html.UnescapeString(rss.Channel.Item[index].Title)
		rss.Channel.Item[index].Description = html.UnescapeString(rss.Channel.Item[index].Description)
	}

	return &rss, nil
}
