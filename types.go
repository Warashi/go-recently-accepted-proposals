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

// acceptedAt returns the date the proposal was accepted.
func (r QueryResult) acceptedAt(i int) time.Time {
	for _, edge := range r.Data.Repository.Issues.Edges[i].Node.TimelineItems.Edges {
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

// QueryResult is the result of the GitHub GraphQL API query.
type QueryResult struct {
	Data struct {
		Repository struct {
			Issues struct {
				Edges []struct {
					Node struct {
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
					} `json:"node"`
				} `json:"edges"`
			} `json:"issues"`
		} `json:"repository"`
	} `json:"data"`
}
