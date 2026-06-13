package main

import (
	"log"
	"os"
	"slices"
)

func main() {
	queryResult, err := parseQueryResult(os.Stdin)
	if err != nil {
		log.Fatalf("Failed to parse query result: %v", err)
	}

	// Sort desc the proposals by the date they were accepted
	slices.SortStableFunc(queryResult.Data.Repository.Issues.Edges, func(a, b Edge) int {
		return a.acceptedAt().Compare(b.acceptedAt())
	})

	// Only show the first 10 proposals
	queryResult.Data.Repository.Issues.Edges = queryResult.Data.Repository.Issues.Edges[:10]

	if err := renderATOM(os.Stdout, queryResult); err != nil {
		log.Fatalf("Failed to render ATOM: %v", err)
	}
}
