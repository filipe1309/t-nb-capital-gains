.PHONY: test-cover coverage test build mod run docker-build docker-run clean help

CLI_BINARY := capital-gains
CLI_MAIN := cmd/cli/capital-gains/main.go
DOCKER_IMAGE := capital-gains:1
COVERAGE_FILE := cp.out
COVERAGE_HTML := coverage.html

all: build docker-build

test-cover:
	go test -race -coverprofile=$(COVERAGE_FILE) ./...

coverage:
	go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)

test: test-cover coverage

build: mod
	@echo "🏗️ Building code..."
	@go build -o bin/$(CLI_BINARY) $(CLI_MAIN)

mod:
	@echo "📦 Downloading dependencies..."
	go mod download

run:
	@go run $(CLI_MAIN)

docker-build:
	@echo "🐳 Building Docker image..."
	@docker build -t $(DOCKER_IMAGE) . --no-cache

docker-run:
	@docker run --rm -i $(DOCKER_IMAGE)

clean:
	@echo "🧹 Cleaning up..."
	@-go clean -testcache
	@-rm -rf bin/$(CLI_BINARY)
	@-docker rmi $(DOCKER_IMAGE)

help:
	@echo "📖 Available commands:"
	@echo "  make test         - Run tests"
	@echo "  make build        - Build code"
	@echo "  make mod          - Download dependencies"
	@echo "  make run          - Run code"
	@echo "  make docker-build - Build Docker image"
	@echo "  make docker-run   - Run Docker container"
	@echo "  make clean        - Clean up"
	@echo "  make help         - Show this help message"
