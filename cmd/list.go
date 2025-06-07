package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [category]",
	Short: "List HTTP status codes",
	Long: `List all HTTP status codes or filter by category.
Categories are: 1xx, 2xx, 3xx, 4xx, 5xx`,
	ValidArgs: []string{"1xx", "2xx", "3xx", "4xx", "5xx"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			listCodes(args[0])
		} else {
			listCodes("")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listCodes(category string) {
	if category == "" {
		// List all codes
		fmt.Println("All HTTP Status Codes:")
		fmt.Println("---------------------")
		
		// Group by category
		for i := 1; i <= 5; i++ {
			fmt.Printf("\n%dxx - ", i)
			switch i {
			case 1:
				fmt.Println("Informational")
			case 2:
				fmt.Println("Success")
			case 3:
				fmt.Println("Redirection")
			case 4:
				fmt.Println("Client Error")
			case 5:
				fmt.Println("Server Error")
			}
			fmt.Println("---------------------")
			
			for code, info := range httpCodesInfo {
				if code/100 == i {
					fmt.Printf("%d: %s\n", code, info.Description)
				}
			}
		}
		return
	}
	
	// List codes by category
	category = strings.ToLower(category)
	if !strings.HasSuffix(category, "xx") || len(category) != 3 {
		fmt.Println("Invalid category. Use 1xx, 2xx, 3xx, 4xx, or 5xx.")
		return
	}
	
	prefix := category[0] - '0'
	if prefix < 1 || prefix > 5 {
		fmt.Println("Invalid category. Use 1xx, 2xx, 3xx, 4xx, or 5xx.")
		return
	}
	
	var categoryName string
	switch prefix {
	case 1:
		categoryName = "Informational"
	case 2:
		categoryName = "Success"
	case 3:
		categoryName = "Redirection"
	case 4:
		categoryName = "Client Error"
	case 5:
		categoryName = "Server Error"
	}
	
	fmt.Printf("%dxx - %s:\n", prefix, categoryName)
	fmt.Println("---------------------")
	
	found := false
	for code, info := range httpCodesInfo {
		if code/100 == int(prefix) {
			fmt.Printf("%d: %s\n", code, info.Description)
			found = true
		}
	}
	
	if !found {
		fmt.Printf("No HTTP status codes found in category %s\n", category)
	}
}
