#!/bin/bash

# Exit script on error
set -e

# Set the project root directory
PROJECT_ROOT=$(pwd)

# Set the output directory
OUTPUT_DIR="$PROJECT_ROOT/bin"
CMD_DIR="$PROJECT_ROOT/cmd"

# Create the output directory if it doesn't exist
mkdir -p "$OUTPUT_DIR"

# Format all go files
echo "Formatting Go files..."
go fmt ./...

# Vet the code (reports suspicious constructs)
echo "Vetting code..."
go vet ./...

# Run tests with coverage
echo "Running tests..."
go test ./... -cover

# Build the Go binary
go build -o "$OUTPUT_DIR/app" "$CMD_DIR/main.go"

# Check if the build was successful
if [ $? -eq 0 ]; then
    echo "Build successful! Binary is located at $OUTPUT_DIR/app"
else
    echo "Build failed!"
fi