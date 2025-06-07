# HTTP Code CLI Tool

A simple command-line tool for looking up HTTP status codes and their descriptions, built with Cobra.

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
httpcode search <term>   - Search for codes by description
httpcode fzf             - Interactive fuzzy search using fzf
httpcode fzf-search      - Interactive fuzzy search with detailed preview
httpcode help            - Show help message
```

## Features

### Detailed Information

For each HTTP status code, the tool provides:
- Short description
- Detailed explanation
- Link to MDN documentation

### Fuzzy Search

The tool includes built-in interactive fuzzy search functionality:

```bash
# Basic fuzzy search
httpcode fzf

# Fuzzy search with detailed preview
httpcode fzf-search
```

The fuzzy search interface allows you to:
- Type to filter HTTP status codes
- Use arrow keys to navigate
- Press Enter to select a code
- View detailed information in the preview pane (with fzf-search)
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
# Look up a specific status code
httpcode 404

# List all 4xx (client error) status codes
httpcode list 4xx

# Search for status codes containing "not found" in their description
httpcode search "not found"

# Interactive fuzzy search
httpcode fzf
```

## License

MIT
