#!/bin/bash

# Breeze Build Script
# Cross-compiles for multiple platforms

echo "Building Breeze for multiple platforms..."

# Linux
echo "Building for Linux..."
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/breeze-linux-amd64 ./cmd/breeze
GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o bin/breeze-linux-arm64 ./cmd/breeze

# macOS
echo "Building for macOS..."
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o bin/breeze-darwin-amd64 ./cmd/breeze
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o bin/breeze-darwin-arm64 ./cmd/breeze

# Windows
echo "Building for Windows..."
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o bin/breeze-windows-amd64.exe ./cmd/breeze
GOOS=windows GOARCH=386 go build -ldflags="-s -w" -o bin/breeze-windows-386.exe ./cmd/breeze

echo "Build complete! Binaries in bin/ directory."