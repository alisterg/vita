# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GORUN = $(GOCMD) run
GOCLEAN = $(GOCMD) clean
GOTIDY = $(GOCMD) mod tidy
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get

# Binary name
BINARY_NAME = vita

all: build

build:
	$(GOBUILD) -tags cli -o $(BINARY_NAME) -v
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)_api -tags api

clean:
	$(GOTIDY)
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_NAME)_api

test:
	$(GOTEST) -v ./...

run:
	$(GORUN) main_cli.go

runapi:
	$(GORUN) main_api.go

.PHONY: all build clean test run
