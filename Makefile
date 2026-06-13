.DELETE_ON_ERROR:
.ONESHELL:

LIMIT ?= 100

.PHONY: all
all: dist/atom.xml

issues.json: query.graphql
	gh api graphql -f query="$$(cat $<)" -F limit=$(LIMIT) > $@

dist/atom.xml: issues.json
	mkdir -p dist
	go run . < $< > $@
