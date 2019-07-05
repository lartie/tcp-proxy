PROJECT_NAME := tcp-proxy
PKG_LIST := $(shell go list ./... | grep -v /vendor/)

DIR := $(shell pwd)

.PHONY: all dep coverage test build clean help

all: build

dep: ## Get the dependencies
	dep ensure -v

coverage: dep ## Run unit tests with coverage
	go test -cover -short $(PKG_LIST)

test: dep ## Run unit tests
	go test -short $(PKG_LIST)

build: dep ## Build the binary file
	env GOOS=linux GOARCH=386 go build -ldflags "-X main.minversion=`date -u +.%Y%m%d.%H%M%S`"

clean: ## Remove previous build
	rm -f $(PROJECT_NAME)

help %: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'