# HTTP Code CLI Tool

A beautiful command-line tool for looking up HTTP status codes and their descriptions, built with Cobra and styled with Lipgloss.

## Features

✨ **Beautiful Terminal UI** - Styled with [Lipgloss](https://github.com/charmbracelet/lipgloss) for a modern, colorful interface
📋 **Comprehensive Database** - Complete HTTP status code information with detailed descriptions
🔍 **Multiple Lookup Options** - Look up by code or browse by category
🎯 **Interactive Fuzzy Search** - Built-in fuzzy search with fzf integration and detailed preview
🌈 **Color-Coded Categories** - Different colors for each HTTP status code category (1xx-5xx)
📖 **MDN Documentation Links** - Direct links to official documentation for each status code

## Installation

First, make sure you have Go installed on your system. Then, you can install the tool by running:

```bash
# Clone the repository
git clone https://github.com/lqt/httpcode.git
cd httpcode

# Build the binary
go build -o httpcode

# Optionally, move the binary to a directory in your PATH
sudo mv httpcode /usr/local/bin/
```

## Usage

```
httpcode <code>          - Look up a specific HTTP status code
httpcode list            - List all HTTP status codes
httpcode list <category> - List codes by category (1xx, 2xx, 3xx, 4xx, 5xx)
httpcode fzf             - Interactive fuzzy search with detailed preview
httpcode help            - Show help message
```

## Beautiful Output

The tool uses Lipgloss to provide beautiful, color-coded output:

- 🔵 **1xx (Informational)** - Blue styling
- 🟢 **2xx (Success)** - Green styling  
- 🟠 **3xx (Redirection)** - Orange styling
- 🔴 **4xx (Client Error)** - Red styling
- 🟣 **5xx (Server Error)** - Purple styling

Each status code is displayed with:
- Styled header with code and description
- Category badge
- Detailed explanation in a bordered box
- Clickable MDN documentation link
- Beautiful separators and formatting

## Detailed Information

For each HTTP status code, the tool provides:
- Short description
- Detailed explanation
- Link to MDN documentation
- Color-coded category classification

## Fuzzy Search

The tool includes built-in interactive fuzzy search functionality:

```bash
# Interactive fuzzy search with detailed preview
httpcode fzf
```

The fuzzy search interface allows you to:
- Type to filter HTTP status codes
- Use arrow keys to navigate
- Press Enter to select a code
- View detailed information in the preview pane
- Press Ctrl+C or Esc to exit

## Shell Completion

The tool supports shell completion for bash, zsh, fish, and PowerShell:

```bash
# Bash
source <(httpcode completion bash)

# Zsh
httpcode completion zsh > "${fpath[1]}/_httpcode"

# Fish
httpcode completion fish > ~/.config/fish/completions/httpcode.fish

# PowerShell
httpcode completion powershell | Out-String | Invoke-Expression
```

To install completions permanently, see the output of `httpcode completion --help`.

## Examples

```bash
# Look up a specific status code with beautiful styling
httpcode 404

# List all 4xx (client error) status codes with color coding
httpcode list 4xx

# Interactive fuzzy search with preview
httpcode fzf

# List all codes with beautiful category headers
httpcode list
```

## Dependencies

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Terminal styling
- [fzf](https://github.com/junegunn/fzf) - Fuzzy search functionality

## Demo

Run the demo script to see the beautiful styling in action:

```bash
chmod +x demo.sh
./demo.sh
```

## License

MIT
