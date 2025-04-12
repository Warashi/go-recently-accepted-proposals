package main

import (
	"encoding/json"
	"io"
	"time"
)

func parseQueryResult(r io.Reader) (QueryResult, error) {
	var result QueryResult
	if err := json.NewDecoder(r).Decode(&result); err != nil {
		return QueryResult{}, err
	}
	return result, nil
}

type QueryResult struct {
	Data struct {
		Repository struct {
			Issues struct {
				Edges []struct {
					Node struct {
						URL           string `json:"url"`
						Number        int    `json:"number"`
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
