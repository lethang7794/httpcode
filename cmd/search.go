package cmd

import (
	"fmt"
	"os"
	"sort"
	"strings"

	fzf "github.com/junegunn/fzf/src"
	"github.com/spf13/cobra"
)


// escapeString escapes special characters in a string for shell command usage
func escapeString(s string) string {
	// Replace characters that could cause issues in shell commands
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "'", "")  // Remove single quotes entirely
	s = strings.ReplaceAll(s, "\"", "")  // Remove double quotes entirely
	s = strings.ReplaceAll(s, "`", "")  // Remove backticks entirely
	s = strings.ReplaceAll(s, "$", "\\$")
	s = strings.ReplaceAll(s, "!", "\\!")
	s = strings.ReplaceAll(s, "&", "\\&")
	s = strings.ReplaceAll(s, "|", "\\|")
	s = strings.ReplaceAll(s, ">", "\\>")
	s = strings.ReplaceAll(s, "<", "\\<")
	s = strings.ReplaceAll(s, "(", "\\(")
	s = strings.ReplaceAll(s, ")", "\\)")
	s = strings.ReplaceAll(s, "[", "\\[")
	s = strings.ReplaceAll(s, "]", "\\]")
	s = strings.ReplaceAll(s, "{", "\\{")
	s = strings.ReplaceAll(s, "}", "\\}")
	s = strings.ReplaceAll(s, ";", "\\;")
	s = strings.ReplaceAll(s, "\n", " ") // Replace newlines with spaces
	return s
}

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Interactive fuzzy search with detailed preview",
	Long:  `Use fuzzy search to interactively search for HTTP status codes with detailed preview.`,
	Run: func(cmd *cobra.Command, args []string) {
		runFzfSearch()
	},
}

func runFzfSearch() {
	// Prepare data for fuzzy search
	var items []string
	var codeMap = make(map[string]int)

	// Sort codes for consistent display
	var codes []int
	for code := range httpCodesInfo {
		codes = append(codes, code)
	}
	sort.Ints(codes)

	// Format items for display with preview information
	for _, code := range codes {
		info := httpCodesInfo[code]
		category := ""
		switch code / 100 {
		case 1:
			category = "Informational"
		case 2:
			category = "Success"
		case 3:
			category = "Redirection"
		case 4:
			category = "Client Error"
		case 5:
			category = "Server Error"
		}

		// Include all information for preview mode
		// Escape special characters in the detail text
		escapedDescription := escapeString(info.Description)
		escapedDetail := escapeString(info.Detail)
		escapedLink := escapeString(info.MDNLink)
		
		item := fmt.Sprintf("%d\t%s\t%s\t%s\t%s", 
			code, 
			escapedDescription, 
			category, 
			escapedDetail, 
			escapedLink)
		
		items = append(items, item)
		codeMap[item] = code
	}

	// Create input channel for fzf
	inputChan := make(chan string)
	go func() {
		for _, item := range items {
			inputChan <- item
		}
		close(inputChan)
	}()

	// Create output channel for fzf results
	outputChan := make(chan string, 1)

	// Exit function
	exit := func(code int, err error) {
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
		// Don't exit the program, just return from the function
		if code != 0 && code != 130 {
			fmt.Fprintln(os.Stderr, "Error in fuzzy search")
		}
	}

	// Build fzf options
	var fzfArgs []string
	
	// Basic options
	fzfArgs = append(fzfArgs, "--ansi", "--reverse", "--border")
	
	// Set height
	fzfArgs = append(fzfArgs, "--height=80%")
	
	// Add header
	fzfArgs = append(fzfArgs, "--header=Code    Message          (Press ESC to exit, Enter to select)")
	
	// Add header label with program information
	fzfArgs = append(fzfArgs, "--border-label=httpcode - HTTP Status Code Viewer")
	
	// Add preview options for detailed view
	previewCmd := "echo -e '\\033[1;32mHTTP Status Code:\\033[0m {1} {2}\\n" +
		"\\033[1;32mClass:\\033[0m            {3}\\n" +
		"\\033[1;32mDetails:\\033[0m\\n{4}\\n" +
		"\\033[1;32mMDN Docs:\\033[0m\\n{5}'"
	
	fzfArgs = append(fzfArgs, 
		"--delimiter=\\t",
		"--with-nth=1,2",
		"--preview=" + previewCmd,
		"--preview-window=right:60%:wrap")

	// Parse options
	options, err := fzf.ParseOptions(false, fzfArgs)
	if err != nil {
		exit(fzf.ExitError, err)
		return
	}

	// Set up input and output channels
	options.Input = inputChan
	options.Output = outputChan

	// Run fzf
	code, err := fzf.Run(options)
	
	// Process the selected item after fzf exits
	select {
	case selection := <-outputChan:
		if statusCode, exists := codeMap[selection]; exists {
			info := httpCodesInfo[statusCode]
			displayCodeWithLipgloss(statusCode, info)
		} else {
			fmt.Println(selection)
		}
	default:
		// No selection made (user cancelled)
	}
	
	exit(code, err)
}

func init() {
	rootCmd.AddCommand(searchCmd)
}