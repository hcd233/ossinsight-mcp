# OSS Insight MCP Server Makefile

# Variables
BINARY_NAME=ossinsight-mcp
BUILD_DIR=build
MAIN_FILE=main.go

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# Build flags
LDFLAGS=-ldflags "-X main.Version=$(shell git describe --tags --always --dirty) -X main.BuildTime=$(shell date -u '+%Y-%m-%d_%H:%M:%S')"

.PHONY: all build clean test deps run server client help

# Default target
all: clean deps test build

# Build the application
build:
	@echo "Building $(BINARY_NAME)..."
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_FILE)
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)
	@echo "Clean complete"

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...
	@echo "Tests complete"

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run tests for specific package
test-package:
	@echo "Running tests for package: $(PACKAGE)"
	$(GOTEST) -v ./$(PACKAGE)

# Install dependencies
deps:
	@echo "Installing dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy
	@echo "Dependencies installed"

# Run the application
run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BUILD_DIR)/$(BINARY_NAME)

# Start the MCP server
server: build
	@echo "Starting MCP server..."
	./$(BUILD_DIR)/$(BINARY_NAME) server --host localhost --port 8080

# Start the MCP client
client: build
	@echo "Starting MCP client..."
	./$(BUILD_DIR)/$(BINARY_NAME) client --endpoint http://localhost:8080/mcp

# Run integration tests
test-integration: build
	@echo "Running integration tests..."
	./$(BUILD_DIR)/$(BINARY_NAME) client --endpoint http://localhost:8080/mcp

# Format code
fmt:
	@echo "Formatting code..."
	$(GOCMD) fmt ./...
	@echo "Code formatting complete"

# Lint code
lint:
	@echo "Linting code..."
	golangci-lint run
	@echo "Linting complete"

# Generate documentation
docs:
	@echo "Generating documentation..."
	$(GOCMD) doc -all ./...
	@echo "Documentation generation complete"

# Create release build
release: clean deps test
	@echo "Creating release build..."
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN_FILE)
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(MAIN_FILE)
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(MAIN_FILE)
	@echo "Release builds complete"

# Install the application
install: build
	@echo "Installing $(BINARY_NAME)..."
	cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/
	@echo "Installation complete"

# Uninstall the application
uninstall:
	@echo "Uninstalling $(BINARY_NAME)..."
	rm -f /usr/local/bin/$(BINARY_NAME)
	@echo "Uninstallation complete"

# Docker build
docker-build:
	@echo "Building Docker image..."
	docker build -t $(BINARY_NAME):latest .
	@echo "Docker build complete"

# Docker run
docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 $(BINARY_NAME):latest

# Show help
help:
	@echo "Available targets:"
	@echo "  all              - Clean, install deps, test, and build"
	@echo "  build            - Build the application"
	@echo "  clean            - Clean build artifacts"
	@echo "  test             - Run all tests"
	@echo "  test-coverage    - Run tests with coverage report"
	@echo "  test-package     - Run tests for specific package (set PACKAGE=path)"
	@echo "  deps             - Install dependencies"
	@echo "  run              - Build and run the application"
	@echo "  server           - Start the MCP server"
	@echo "  client           - Start the MCP client"
	@echo "  test-integration - Run integration tests"
	@echo "  fmt              - Format code"
	@echo "  lint             - Lint code (requires golangci-lint)"
	@echo "  docs             - Generate documentation"
	@echo "  release          - Create release builds for multiple platforms"
	@echo "  install          - Install the application"
	@echo "  uninstall        - Uninstall the application"
	@echo "  docker-build     - Build Docker image"
	@echo "  docker-run       - Run Docker container"
	@echo "  help             - Show this help message"

# Development helpers
dev-setup: deps fmt lint
	@echo "Development environment setup complete"

# Quick test
quick-test:
	@echo "Running quick tests..."
	$(GOTEST) -v ./internal/repo/ossinsight
	$(GOTEST) -v ./internal/repo/mcp
	@echo "Quick tests complete"

# Benchmark tests
bench:
	@echo "Running benchmark tests..."
	$(GOTEST) -bench=. ./...
	@echo "Benchmark tests complete"

# Race condition tests
race:
	@echo "Running race condition tests..."
	$(GOTEST) -race ./...
	@echo "Race condition tests complete"

# Vendor dependencies
vendor:
	@echo "Vendoring dependencies..."
	$(GOMOD) vendor
	@echo "Dependencies vendored"

# Update dependencies
update-deps:
	@echo "Updating dependencies..."
	$(GOGET) -u ./...
	$(GOMOD) tidy
	@echo "Dependencies updated"