.PHONY: help test run build docker-build docker-run clean

CLI_BINARY := bin/capital-gains
CLI_MAIN := cmd/cli/capital-gains/main.go
DOCKER_IMAGE := capital-gains:latest

all: build docker-build

test:
	@echo "🟢 Running tests..."
	go test -v ./...

build:
	@echo "🏗️ Building code..."
	go build -o $(CLI_BINARY) $(CLI_MAIN)

run:
	go run $(CLI_MAIN)

docker-build:
	@echo "🐳 Building Docker image..."
	@docker build -t $(DOCKER_IMAGE) . ||:

docker-run:
	@docker run --rm -i $(DOCKER_IMAGE)

clean:
	@echo "🧹 Cleaning up..."
	@-go clean -testcache
	@-rm -rf $(CLI_BINARY)
	@-docker rmi $(DOCKER_IMAGE)

help:
	@echo "📖 Available commands:"
	@echo "  make test         - Run tests"
	@echo "  make build        - Build code"
	@echo "  make run          - Run code"
	@echo "  make docker-build - Build Docker image"
	@echo "  make docker-run   - Run Docker container"
	@echo "  make clean        - Clean up"
	@echo "  make help         - Show this help message"
