package internal

import (
	"github.com/gorilla/feeds"
	"time"
)

func generateFeed() (string, error) {
	now := time.Now()

	feed := &feeds.Feed{
		Title:       "Omnivore RSS",
		Link:        &feeds.Link{Href: "https://omnivore.app"},
		Description: "Automatically generated RSS feed based on the last items in your Omnivore account",
		Created: now,
	}

	items := []*feeds.Item{}
	for _, item := range SearchQuery.Search.SearchSuccess.Edges {

		items = append(items, &feeds.Item{
			Id: item.Node.Id,
			Title:   item.Node.Title,
			Created: item.Node.CreatedAt,
			Author: &feeds.Author{
				Name:  item.Node.Author,
				Email: item.Node.Author,
			},
			Link: &feeds.Link{
				Href: item.Node.Url,
			},
			Description: item.Node.Description,
			Content: item.Node.Content,
		})
	}

	feed.Items = items

	return feed.ToRss()
}
