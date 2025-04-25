#!/bin/bash

set -e  # Exit on any error

echo "Running go tests..."
go test ./... -v

echo "All checks passed. Starting the application..."
go run  cmd/main.go