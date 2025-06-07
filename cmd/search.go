package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search <term>",
	Short: "Search for HTTP status codes by description",
	Long:  `Search for HTTP status codes containing the specified term in their description.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		searchCodes(args[0])
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}

func searchCodes(term string) {
	term = strings.ToLower(term)
	var matchedCodes []int
	
	displayListHeaderWithLipgloss(fmt.Sprintf("Search Results for \"%s\"", term))
	
	// Find matching codes
	for code, info := range httpCodesInfo {
		searchText := strings.ToLower(info.Description + " " + info.Detail)
		if strings.Contains(searchText, term) {
			matchedCodes = append(matchedCodes, code)
		}
	}
	
	if len(matchedCodes) == 0 {
		displayErrorWithLipgloss(fmt.Sprintf("No HTTP status codes found matching \"%s\"", term))
		return
	}
	
	// Sort matched codes
	sort.Ints(matchedCodes)
	
	// Display search results with enhanced styling
	for _, code := range matchedCodes {
		info := httpCodesInfo[code]
		detail := truncateText(info.Detail, 100)
		displaySearchResultWithLipgloss(code, info, detail)
	}
	
	// Display summary
	displaySummaryWithLipgloss(fmt.Sprintf("Found %d result(s)", len(matchedCodes)))
}

// Helper function to truncate text with ellipsis if it's too long
func truncateText(text string, maxLen int) string {
	if len(text) <= maxLen {
		return text
	}
	return text[:maxLen-3] + "..."
}
