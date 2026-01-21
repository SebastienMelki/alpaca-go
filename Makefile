.PHONY: all generate build lint lint-fix buf-lint test clean release install-tools deps check

# Default target
all: generate build

# Install required tools
install-tools:
	go install github.com/SebastienMelki/sebuf/cmd/protoc-gen-go-client@latest
	go install github.com/SebastienMelki/sebuf/cmd/protoc-gen-openapiv3@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Update buf dependencies
deps:
	buf dep update

# Generate Go code and OpenAPI specs from proto files
generate:
	buf generate

# Build the Go code
build:
	go build ./...

# Run Go linter
lint:
	@echo "Running linter..."
	@golangci-lint run ./...

# Run Go linter with auto-fix
lint-fix:
	@echo "Running linter with auto-fix..."
	@golangci-lint run --fix ./...

# Run buf lint on proto files
buf-lint:
	buf lint

# Run tests
test:
	go test ./...

# Clean generated files
clean:
	rm -rf internal/gen/ docs/

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

# Run all checks (buf-lint, lint, generate, build, test)
check: buf-lint lint generate build test
