package main

import (
	"io"
	"strconv"
	"time"

	"github.com/gorilla/feeds"
)

func renderATOM(dst io.Writer, queryResult []QueryResult) error {
	f := &feeds.Feed{
		Title:       "Go's Recently Accepted Proposals",
		Link:        &feeds.Link{Href: "https://github.com/golang/go/issues?q=is%3Aissue%20state%3Aopen%20label%3AProposal-Accepted"},
		Description: "Go's Recently Accepted Proposals",
		Created:     time.Now(),
	}

	for _, r := range queryResult {
		issue := r.Data.Repository.Issue
		f.Items = append(f.Items, &feeds.Item{
			Title:       issue.Title + " #" + strconv.Itoa(issue.Number),
			Link:        &feeds.Link{Href: issue.URL},
			Description: issue.BodyText,
			Created:     issue.acceptedAt(),
		})
	}

	return f.WriteAtom(dst)
}
