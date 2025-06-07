package cmd

import (
	"fmt"
	"os"
	"sort"

	fzf "github.com/junegunn/fzf/src"
	"github.com/spf13/cobra"
)

// fzfCmd represents the fzf command
var fzfCmd = &cobra.Command{
	Use:   "fzf",
	Short: "Interactive fuzzy search for HTTP status codes",
	Long:  `Use fuzzy search to interactively search for HTTP status codes.`,
	Run: func(cmd *cobra.Command, args []string) {
		runFzfSearch(false)
	},
}

// fzfSearchCmd represents the fzf-search command
var fzfSearchCmd = &cobra.Command{
	Use:   "fzf-search",
	Short: "Interactive fuzzy search with detailed preview",
	Long:  `Use fuzzy search to interactively search for HTTP status codes with detailed preview.`,
	Run: func(cmd *cobra.Command, args []string) {
		runFzfSearch(true)
	},
}

func runFzfSearch(withPreview bool) {
	// Prepare data for fuzzy search
	var items []string
	var codeMap = make(map[string]int)

	// Sort codes for consistent display
	var codes []int
	for code := range httpCodesInfo {
		codes = append(codes, code)
	}
	sort.Ints(codes)

	// Format items for display
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

		var item string
		if withPreview {
			// Include all information for preview mode
			item = fmt.Sprintf("%d\t%s\t%s\t%s\t%s", 
				code, 
				info.Description, 
				category, 
				info.Detail, 
				info.MDNLink)
		} else {
			item = fmt.Sprintf("%d: %s", code, info.Description)
		}
		
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
	outputChan := make(chan string)
	go func() {
		for selection := range outputChan {
			if code, exists := codeMap[selection]; exists {
				info := httpCodesInfo[code]
				fmt.Printf("%d: %s\n", code, info.Description)
				fmt.Printf("\nDetail: %s\n", info.Detail)
				fmt.Printf("\nMDN Documentation: %s\n", info.MDNLink)
			} else {
				// This should not happen, but just in case
				fmt.Println(selection)
			}
		}
	}()

	// Exit function
	exit := func(code int, err error) {
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
		// Don't exit the program, just return from the function
		if code != 0 {
			fmt.Fprintln(os.Stderr, "Error in fuzzy search")
		}
	}

	// Build fzf options
	var fzfArgs []string
	
	// Basic options
	fzfArgs = append(fzfArgs, "--ansi", "--reverse", "--border")
	
	// Set height
	fzfArgs = append(fzfArgs, "--height=50%")
	
	// Add header
	fzfArgs = append(fzfArgs, "--header=HTTP Status Codes (Press ESC to exit, Enter to select)")
	
	if withPreview {
		// Add preview options for detailed view
		previewCmd := "echo -e '\\033[1;32mHTTP Status Code:\\033[0m {1}\\n\\n" +
			"\\033[1;32mDescription:\\033[0m {2}\\n\\n" +
			"\\033[1;32mCategory:\\033[0m {3}\\n\\n" +
			"\\033[1;32mDetails:\\033[0m\\n{4}\\n\\n" +
			"\\033[1;32mMDN Documentation:\\033[0m\\n{5}'"
		
		fzfArgs = append(fzfArgs, 
			"--delimiter=\\t",
			"--with-nth=1,2",
			"--preview=" + previewCmd,
			"--preview-window=right:50%:wrap")
	}

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
	exit(code, err)
}

func init() {
	rootCmd.AddCommand(fzfCmd)
	rootCmd.AddCommand(fzfSearchCmd)
}