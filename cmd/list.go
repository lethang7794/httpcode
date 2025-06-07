package cmd

import (
	"fmt"
	"sort"
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
		displayListHeaderWithLipgloss("All HTTP Status Codes")
		
		// Group by category
		for i := 1; i <= 5; i++ {
			var categoryName string
			switch i {
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
			
			// Display category header
			displayCategoryHeaderWithLipgloss(i, categoryName)
			
			// Get and sort codes for this category
			var codes []int
			for code := range httpCodesInfo {
				if code/100 == i {
					codes = append(codes, code)
				}
			}
			sort.Ints(codes)
			
			// Display codes in this category
			for _, code := range codes {
				displayCodeListItemWithLipgloss(code, httpCodesInfo[code].Description)
			}
		}
		return
	}
	
	// List codes by category
	category = strings.ToLower(category)
	if !strings.HasSuffix(category, "xx") || len(category) != 3 {
		displayErrorWithLipgloss("Invalid category. Use 1xx, 2xx, 3xx, 4xx, or 5xx.")
		return
	}
	
	prefix := category[0] - '0'
	if prefix < 1 || prefix > 5 {
		displayErrorWithLipgloss("Invalid category. Use 1xx, 2xx, 3xx, 4xx, or 5xx.")
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
	
	displayListHeaderWithLipgloss(fmt.Sprintf("%dxx - %s", prefix, categoryName))
	
	// Get and sort codes for this category
	var codes []int
	for code := range httpCodesInfo {
		if code/100 == int(prefix) {
			codes = append(codes, code)
		}
	}
	
	if len(codes) == 0 {
		displayErrorWithLipgloss(fmt.Sprintf("No HTTP status codes found in category %s", category))
		return
	}
	
	sort.Ints(codes)
	
	// Display codes in this category
	for _, code := range codes {
		displayCodeListItemWithLipgloss(code, httpCodesInfo[code].Description)
	}
}
