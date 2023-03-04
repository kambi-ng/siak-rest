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

doc-docker:
	$(docker inspect --type=image swaggo/swag &> /dev/null || docker build -t swaggo/swag -f tools/Dockerfile.swag .) 
	docker run --rm -it -v $(PWD):/app -w /app swaggo/swag sh -c "swag init -g cmd/http-server/main.go"

tool:
	@sh .tools/install.sh

.PHONY: build doc doc-docker dev install run run-build tool
