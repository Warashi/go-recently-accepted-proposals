.DELETE_ON_ERROR:
.ONESHELL:

LIMIT ?= 100

.PHONY: all
all: dist/atom.xml

build/issues.json: query.graphql
	mkdir -p build
	gh api graphql -f query="$$(cat $<)" -F limit=$(LIMIT) > $@

dist/atom.xml: build/issues.json
	mkdir -p dist
	go run . < $< > $@
