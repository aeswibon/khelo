# Constants

ifeq ($(OS),Windows_NT) 
    DETECTED_OS := Windows
else
    DETECTED_OS := $(shell sh -c 'uname 2>/dev/null || echo Unknown')
endif

# Help

.SILENT: help
help:
	@echo Usage: make [command]
	@echo Commands:
	@echo build                         Build khelo server
	@echo run                           Run khelo server
	@echo docker-up                     Up docker services
	@echo docker-down                   Down docker services
	@echo fmt                           Format source code
	@echo test                          Run unit tests

# Build

.SILENT: build
build:
	@echo Build
	@air build

# Run

.SILENT: run
run:
	@air run


# Docker

.SILENT: docker-up
docker-up:
	@docker-compose up -d

.SILENT: docker-down
docker-down:
	@docker-compose down


# Format

.SILENT: fmt
fmt:
	@go fmt ./...

# Default
.DEFAULT_GOAL := help
