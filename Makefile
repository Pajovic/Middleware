MAKEFLAGS += --silent
SHELL := /bin/bash
BIN_NAME := middleware

#Project related commands
build:
	echo "Compiling $(BIN_NAME)..."
	go build -o $(BIN_NAME)
	
.PHONY: build

clean: 
	rm -f $(BIN_NAME)

.PHONY: clean

test:
	echo "Starting tests..."
	go test ./...

.PHONY: test

#Docker commands
deploy:
	echo "Deploying docker image $(BIN_NAME)"
	docker build . -t $(BIN_NAME)

.PHONY: deploy

run:
	echo "Starting docker container $(BIN_NAME)..."
	docker run -d --name $(BIN_NAME) -p 8080:8080 $(BIN_NAME)

.PHONY: run