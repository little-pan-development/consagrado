#!/bin/bash

.SILENT:
.PHONY: help

## dev | Create and start containers
dev-up:
	@docker-compose -f docker/development/docker-compose.yml up -d ${service} --build

## dev | Follow containers logs
dev-logs:
	@docker-compose -f docker/development/docker-compose.yml logs -f ${service}

## dev | Stop and remove containers
dev-down:
	@docker-compose -f docker/development/docker-compose.yml down

## dev | Run tests
dev-test:
	@docker-compose -f docker/development/docker-compose.yml exec palmirinha-app go test -v -cover ./...

## dev | Get cover test
dev-test-cover:
	@docker-compose -f docker/development/docker-compose.yml exec palmirinha-app go test -cover ./... -coverprofile=c.out
	@docker-compose -f docker/development/docker-compose.yml exec palmirinha-app go tool cover -html=c.out -o coverage.html


## Show this help
help:
	printf "Commands:\n"
	awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "%-15s %s\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)