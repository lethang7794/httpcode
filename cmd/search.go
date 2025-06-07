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
	
	for code, desc := range httpCodes {
		if strings.Contains(strings.ToLower(desc), term) {
			fmt.Printf("%d: %s\n", code, desc)
			found = true
		}
	}
	
	if !found {
		fmt.Printf("No HTTP status codes found matching \"%s\"\n", term)
	}
}
