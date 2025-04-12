.DELETE_ON_ERROR:
.ONESHELL:

LIMIT ?= 1

.PHONY: all
all:

issues.json: query.graphql
	mkdir -p src
	gh api graphql -f query="$$(cat $<)" -F limit=$(LIMIT) > $@
