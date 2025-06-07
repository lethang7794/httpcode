package cmd

import (
	"fmt"
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
	found := false
	
	fmt.Printf("Search results for \"%s\":\n", term)
	fmt.Println("---------------------")
	
	for code, info := range httpCodesInfo {
		searchText := strings.ToLower(info.Description + " " + info.Detail)
		if strings.Contains(searchText, term) {
			fmt.Printf("%d: %s\n", code, info.Description)
			fmt.Printf("   %s\n", truncateText(info.Detail, 80))
			fmt.Printf("   %s\n\n", info.MDNLink)
			found = true
		}
	}
	
	if !found {
		fmt.Printf("No HTTP status codes found matching \"%s\"\n", term)
	}
}

// Helper function to truncate text with ellipsis if it's too long
func truncateText(text string, maxLen int) string {
	if len(text) <= maxLen {
		return text
	}
	return text[:maxLen-3] + "..."
}
