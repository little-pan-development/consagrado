#!/bin/bash
.SILENT:
.PHONY: help

include docker/development/env
export

## dev | Create and start containers
dev-up:
	@docker-compose -f docker/development/docker-compose.yml up -d ${service} --build

## dev | Follow containers logs
dev-logs:
	@docker-compose -f docker/development/docker-compose.yml logs -f ${service}

## dev | Install migrations
dev-migration:
	@docker cp docker/development/database.sql "${MYSQL_HOST}":/
	@docker exec -it "${MYSQL_HOST}" bash -c "mysql -u ${MYSQL_USERNAME} -p${MYSQL_PASSWORD} ${MYSQL_DATABASE} < /database.sql"

## dev | Stop and remove containers
dev-down:
	@docker-compose -f docker/development/docker-compose.yml down

## dev | Run tests
dev-test:
	@docker-compose -f docker/development/docker-compose.yml exec ${APPLICATION_NAME} go test -v -cover ./...

## dev | Get cover test
dev-test-cover:
	@docker-compose -f docker/development/docker-compose.yml exec ${APPLICATION_NAME} go test -cover ./... -coverprofile=c.out
	@docker-compose -f docker/development/docker-compose.yml exec ${APPLICATION_NAME} go tool cover -html=c.out -o coverage.html


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