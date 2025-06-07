#!/bin/bash

# Demo script to showcase the beautiful Lipgloss styling

echo "🎨 HTTP Code CLI Tool - Lipgloss Styling Demo"
echo "=============================================="
echo

# Build the tool first
echo "Building httpcode..."
go build -o httpcode
echo

# Demo individual code lookup
echo "📋 Looking up HTTP 404:"
./httpcode 404
echo

echo "📋 Looking up HTTP 200:"
./httpcode 200
echo

echo "📋 Looking up HTTP 500:"
./httpcode 500
echo

# Demo list functionality
echo "📋 Listing 4xx codes:"
./httpcode list 4xx
echo

# Demo fuzzy search functionality
echo "📋 Fuzzy search available with:"
echo "   ./httpcode fzf"
echo

echo "✨ Demo complete! Your HTTP code tool now has beautiful styling with Lipgloss!"
