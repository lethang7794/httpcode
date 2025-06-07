#!/bin/bash

# Script to install fzf if not already installed

# Check if fzf is already installed
if command -v fzf &> /dev/null; then
    echo "fzf is already installed"
    exit 0
fi

echo "Installing fzf..."

# Determine the package manager and install fzf
if command -v apt-get &> /dev/null; then
    # Debian/Ubuntu
    sudo apt-get update
    sudo apt-get install -y fzf
elif command -v dnf &> /dev/null; then
    # Fedora
    sudo dnf install -y fzf
elif command -v yum &> /dev/null; then
    # CentOS/RHEL
    sudo yum install -y fzf
elif command -v pacman &> /dev/null; then
    # Arch Linux
    sudo pacman -S --noconfirm fzf
elif command -v brew &> /dev/null; then
    # macOS with Homebrew
    brew install fzf
else
    # Manual installation using git
    echo "No package manager found. Installing fzf using git..."
    
    if ! command -v git &> /dev/null; then
        echo "Error: git is not installed. Please install git first."
        exit 1
    fi
    
    git clone --depth 1 https://github.com/junegunn/fzf.git ~/.fzf
    ~/.fzf/install --all
fi

# Check if installation was successful
if command -v fzf &> /dev/null; then
    echo "fzf has been successfully installed"
else
    echo "Failed to install fzf"
    echo "Please install fzf manually: https://github.com/junegunn/fzf#installation"
    exit 1
fi
