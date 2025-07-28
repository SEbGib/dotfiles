.PHONY: build run clean install test

# Build the TUI application
build:
	go build -o dotfiles-tui ./cmd/dotfiles-tui

# Run the TUI application
run: build
	./dotfiles-tui

# Clean build artifacts
clean:
	rm -f dotfiles-tui

# Install dependencies
deps:
	go mod tidy
	go mod download

# Install the binary to system
install: build
	sudo mv dotfiles-tui /usr/local/bin/

# Run tests
test:
	go test ./...

# Development mode with hot reload (requires air)
dev:
	air -c .air.toml

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run

# Show help
help:
	@echo "Available commands:"
	@echo "  build    - Build the TUI application"
	@echo "  run      - Build and run the TUI application"
	@echo "  clean    - Clean build artifacts"
	@echo "  deps     - Install dependencies"
	@echo "  install  - Install binary to /usr/local/bin"
	@echo "  test     - Run tests"
	@echo "  dev      - Run in development mode with hot reload"
	@echo "  fmt      - Format code"
	@echo "  lint     - Lint code"
	@echo "  help     - Show this help message"