package main

import (
	"io"
	"strconv"
	"time"

	"github.com/gorilla/feeds"
)

func renderATOM(dst io.Writer, queryResult QueryResult) error {
	f := &feeds.Feed{
		Title:       "Go's Recently Accepted Proposals",
		Link:        &feeds.Link{Href: "https://github.com/golang/go/issues?q=is%3Aissue%20state%3Aopen%20label%3AProposal-Accepted"},
		Description: "Go's Recently Accepted Proposals",
		Created:     time.Now(),
	}

	for i, edge := range queryResult.Data.Repository.Issues.Edges {
		f.Items = append(f.Items, &feeds.Item{
			Title:       edge.Node.Title,
			Link:        &feeds.Link{Href: edge.Node.URL},
			Description: edge.Node.Title + " #" + strconv.Itoa(edge.Node.Number),
			Created:     queryResult.acceptedAt(i),
		})
	}

	return f.WriteAtom(dst)
}
