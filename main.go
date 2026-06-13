package main

import (
	"encoding/json"
	"io/fs"
	"iter"
	"log"
	"os"
	"path/filepath"
	"slices"
)

func walk(root string) iter.Seq2[string, fs.DirEntry] {
	return func(yield func(string, fs.DirEntry) bool) {
		filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return filepath.SkipAll
			}
			if !yield(path, d) {
				return filepath.SkipAll
			}
			return nil
		})
	}
}

func readQueryResult(path string) (QueryResult, error) {
	var queryResult QueryResult
	file, err := os.Open(path)
	if err != nil {
		return queryResult, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&queryResult); err != nil {
		return queryResult, err
	}
	return queryResult, nil
}

func main() {
	results := make([]QueryResult, 0, 100)

	for path, entry := range walk("./build/timeline") {
		if entry.IsDir() {
			continue
		}
		queryResult, err := readQueryResult(path)
		if err != nil {
			log.Printf("Failed to read query result from %s: %v", path, err)
			os.Exit(1)
		}
		results = append(results, queryResult)
	}

	// Sort desc the proposals by the date they were accepted
	slices.SortStableFunc(results, func(a, b QueryResult) int {
		return -a.acceptedAt().Compare(b.acceptedAt())
	})

	if err := renderATOM(os.Stdout, results[:min(10, len(results))]); err != nil {
		log.Fatalf("Failed to render ATOM: %v", err)
	}
}
