.DELETE_ON_ERROR:
.ONESHELL:

LIMIT ?= 100

.PHONY: all
all: atom.xml

issues.json: query.graphql
	mkdir -p src
	gh api graphql -f query="$$(cat $<)" -F limit=$(LIMIT) > $@

atom.xml: issues.json
	go run . < $< > $@
