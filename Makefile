.PHONY: all generate build lint test clean release install-tools deps check

# Default target
all: generate build

# Install required tools
install-tools:
	go install github.com/SebastienMelki/sebuf/cmd/protoc-gen-go-http@latest
	go install github.com/SebastienMelki/sebuf/cmd/protoc-gen-openapiv3@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Update buf dependencies
deps:
	buf dep update

# Generate Go code and OpenAPI specs from proto files
generate:
	buf generate

# Build the Go code
build:
	go build ./...

# Run buf lint on proto files
lint:
	buf lint

# Run tests
test:
	go test ./...

# Clean generated files
clean:
	rm -rf api/ docs/

# Release - creates a new git tag and pushes it
# Usage: make release VERSION=v1.0.0
release:
ifndef VERSION
	$(error VERSION is required. Usage: make release VERSION=v1.0.0)
endif
	@echo "Creating release $(VERSION)..."
	git tag -a $(VERSION) -m "Release $(VERSION)"
	git push origin $(VERSION)
	@echo "Release $(VERSION) created and pushed."

# Run all checks (lint, generate, build, test)
check: lint generate build test
