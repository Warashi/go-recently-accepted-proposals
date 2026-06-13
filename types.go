package main

import (
	"encoding/json"
	"io"
	"time"
)

// parseQueryResult parses the query result from the GitHub GraphQL API.
func parseQueryResult(r io.Reader) (QueryResult, error) {
	var result QueryResult
	if err := json.NewDecoder(r).Decode(&result); err != nil {
		return QueryResult{}, err
	}
	return result, nil
}

// QueryResult is the result of the GitHub GraphQL API query.
type QueryResult struct {
	Data struct {
		Repository struct {
			Issue Issue `json:"issue"`
		} `json:"repository"`
	} `json:"data"`
}

func (q QueryResult) acceptedAt() time.Time {
	return q.Data.Repository.Issue.acceptedAt()
}

type Issue struct {
	Title         string `json:"title"`
	URL           string `json:"url"`
	Number        int    `json:"number"`
	BodyText      string `json:"bodyText"`
	TimelineItems struct {
		Edges []struct {
			Node struct {
				Typename  string    `json:"__typename"`
				CreatedAt time.Time `json:"createdAt"`
				Label     struct {
					Name string `json:"name"`
				} `json:"label"`
			} `json:"node"`
		} `json:"edges"`
	} `json:"timelineItems"`
}

func (i Issue) acceptedAt() time.Time {
	for _, edge := range i.TimelineItems.Edges {
		if edge.Node.Typename != "LabeledEvent" {
			continue
		}
		if edge.Node.Label.Name != "Proposal-Accepted" {
			continue
		}
		return edge.Node.CreatedAt
	}
	return time.Time{}
}
