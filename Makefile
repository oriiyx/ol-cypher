#!/bin/zsh

install:
	@echo "Started installing process."
	@go mod tidy
	@go install
	@echo "Finished installed."

build:
	@echo "Started building process."
	@go build -o bin/ol-cypher
	@echo "Finished building."

build-parser:
	@echo "Started building parser process."
	@go build -o bin/parser cmd/parser/main.go
	@echo "Finished building parser."

test:
	@echo "Started testing process."
	@go mod tidy
	@go install
	@go test -v ./...
	@echo "Finished testing."