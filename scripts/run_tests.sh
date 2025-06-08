#!/bin/bash

# Test runner script for HTTP Code CLI Tool

echo "🧪 Running HTTP Code CLI Tool Tests"
echo "===================================="
echo

# Set test environment
export GO_ENV=test

# Run all tests with verbose output
echo "📋 Running all tests..."
go test -v ./cmd/...

echo
echo "📊 Running tests with coverage..."
go test -v -cover ./cmd/...

echo
echo "📈 Generating detailed coverage report..."
go test -coverprofile=coverage.out ./cmd/...
go tool cover -html=coverage.out -o coverage.html

echo
echo "🔍 Running race condition tests..."
go test -race ./cmd/...

echo
echo "🚀 Running benchmarks..."
go test -bench=. ./cmd/...

echo
echo "✅ Test Summary:"
echo "- Unit tests completed"
echo "- Coverage report generated: coverage.html"
echo "- Race condition tests completed"
echo "- Benchmarks completed"
echo
echo "📁 Test files:"
echo "- cmd/root_test.go      - Root command and lookup tests"
echo "- cmd/list_test.go      - List command tests"
echo "- cmd/search_test.go    - Search command tests"
echo "- cmd/display_test.go   - Display/styling tests"
echo "- cmd/codes_test.go     - HTTP codes data tests"
echo
echo "🎉 All tests completed!"
