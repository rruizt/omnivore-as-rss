package internal

import (
	"context"
	"net/http"
	"time"
	"github.com/hasura/go-graphql-client"
)

var client *graphql.Client

func queryOmnivore() error {
	if client == nil {
		client = graphql.NewClient("https://api-prod.omnivore.app/api/graphql", nil).WithRequestModifier(func(r *http.Request) {
			r.Header.Set("Authorization", Cfg.OmnivoreAuthToken)
		})
	}

	err := client.Query(context.Background(), &SearchQuery, nil)
	if err != nil {
		return err
	}

	return nil
}

var SearchQuery struct {
	Search struct {
		SearchSuccess struct {
			PageInfo struct {
				TotalCount int
			}
			Edges []struct {
				Node struct {
					Id          string
					Title       string
					CreatedAt   time.Time
					Author      string
					Description string
					Url         string
					Image       string
					Content     string
				}
			}
		} `graphql:"... on SearchSuccess"`
		SearchError struct {
			ErrorCodes []string
		} `graphql:"... on SearchError"`
	} `graphql:"search(first: 10, includeContent: true)"`
}
