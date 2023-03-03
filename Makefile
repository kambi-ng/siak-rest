install:
	@go mod tidy

run:
	@go run cmd/http-server/main.go

dev:
	@air

build:
	@go build -o ./bin/main cmd/http-server/main.go

run-build:
	@./bin/main

doc:
	swag init -g cmd/http-server/main.go

tool:
	@sh ./tool.sh

.PHONY: build
