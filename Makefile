.PHONY: help build run test fmt vet lint cross clean

help:
	@echo "Common targets:"
	@echo "  make build   - Build CLI for current platform"
	@echo "  make test    - Run tests"
	@echo "  make fmt     - Format code (go fmt)"
	@echo "  make vet     - Static analysis (go vet)"
	@echo "  make lint    - Vet + optional golangci-lint if installed"
	@echo "  make cross   - Cross-compile via build.sh"
	@echo "  make clean   - Remove local binaries and bin/ artifacts"

build:
	go build ./cmd/breeze

# Usage: make run ARGS='chat "Hello"'
run: build
	./breeze $(ARGS)

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	@command -v golangci-lint >/dev/null 2>&1 && golangci-lint run || echo "golangci-lint not found; running 'go vet' only"
	go vet ./...

cross:
	bash build.sh

clean:
	rm -f breeze
	rm -rf bin/

