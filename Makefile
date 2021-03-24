MAKEFLAGS += --silent
SHELL := /bin/bash
BIN_NAME := middleware


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