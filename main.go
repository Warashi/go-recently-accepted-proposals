package main

import (
	"log"
	"os"
	"sort"
)

func main() {
	queryResult, err := parseQueryResult(os.Stdin)
	if err != nil {
		log.Fatalf("Failed to parse query result: %v", err)
	}

	// Sort the proposals by the date they were accepted
	sort.SliceStable(queryResult.Data.Repository.Issues.Edges, func(i, j int) bool {
		return queryResult.acceptedAt(i).After(queryResult.acceptedAt(j))
	})

	// Only show the first 10 proposals
	queryResult.Data.Repository.Issues.Edges = queryResult.Data.Repository.Issues.Edges[:10]

	if err := renderATOM(os.Stdout, queryResult); err != nil {
		log.Fatalf("Failed to render ATOM: %v", err)
	}
}
