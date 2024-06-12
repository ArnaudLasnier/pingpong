MAKEFLAGS += --silent
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-variables
MAKEFLAGS += --no-builtin-rules


SHELL       := bash
.SHELLFLAGS := -eu -o pipefail -c
.ONESHELL:


.DEFAULT_GOAL := build


NAME ?= pingpong


.PHONY: help
help: ## Display help.
	grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


.PHONY: build
build: ## Build.
	go build -o ./bin/$(NAME) ./cmd/pingpong
	chmod +x ./bin/$(NAME)


.PHONY: run
run: ## Build and run.
	go build -o ./bin/$(NAME) ./cmd/pingpong
	chmod +x ./bin/$(NAME)
	./bin/$(NAME)


.PHONY: live_reload
live_reload: ## Build and run.
	air


.PHONY: dump_schema
dump_schema: ## Dump the database schema.
	pg_dump -d $(DATABASE_URI) --schema-only -f schema.sql


.PHONY: localdeps
localdeps: ## Run the local dependencies and update the local configuration for the app.
	go build -o ./bin/localdeps ./cmd/localdeps
	chmod +x ./bin/localdeps
	./bin/localdeps --env-file ./.envrc


.PHONY: dbgen
dbgen: ## Generate the database models.
	go build -o ./bin/dbgen ./cmd/dbgen
	chmod +x ./bin/dbgen
	./bin/dbgen
