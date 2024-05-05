MAKEFLAGS += --silent
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-variables
MAKEFLAGS += --no-builtin-rules


SHELL       := bash
.SHELLFLAGS := -eu -o pipefail -c
.ONESHELL:


.DEFAULT_GOAL := build


NAME_PROD ?= pingpong
NAME_DEV  ?= pingpongdev


.PHONY: help
help: ## Display help.
	grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


.PHONY: build_prod
build_prod: ## Build the production version.
	go build -o ./bin/$(NAME_PROD) ./cmd/pingpong
	chmod +x ./bin/$(NAME_PROD)


.PHONY: run_dev
run_dev: ## Run the development version.
	go build -tags dev -o ./bin/$(NAME_DEV) ./cmd/pingpong
	chmod +x ./bin/$(NAME_DEV)
	./bin/$(NAME_DEV)


.PHONY: dump_schema
dump_schema: ## Dump the database schema.
	pg_dump -d postgres://postgres:password@localhost:60383/pingpongdb?sslmode=disable --schema-only -f schema.sql
