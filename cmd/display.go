package cmd

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Color palette for different HTTP status code categories
var (
	// Status code colors
	informationalColor = lipgloss.Color("#3498db") // Blue for 1xx
	successColor       = lipgloss.Color("#2ecc71") // Green for 2xx
	redirectionColor   = lipgloss.Color("#f39c12") // Orange for 3xx
	clientErrorColor   = lipgloss.Color("#e74c3c") // Red for 4xx
	serverErrorColor   = lipgloss.Color("#9b59b6") // Purple for 5xx
	unknownColor       = lipgloss.Color("#95a5a6") // Gray for unknown
	
	// UI colors
	textColor      = lipgloss.Color("#2c3e50")
	mutedColor     = lipgloss.Color("#7f8c8d")
	linkColor      = lipgloss.Color("#3498db")
	backgroundColor = lipgloss.Color("#ecf0f1")
	whiteColor     = lipgloss.Color("#ffffff")
)

// Styles
var (
	// Main header style for status codes
	headerStyle = lipgloss.NewStyle().
		Bold(true).
		Padding(1, 2).
		Margin(1, 0).
		Border(lipgloss.DoubleBorder()).
		Align(lipgloss.Center).
		Width(60)
	
	// Category badge style
	badgeStyle = lipgloss.NewStyle().
		Bold(true).
		Padding(0, 2).
		Margin(0, 0, 1, 0).
		Foreground(whiteColor)
	
	// Description box style
	descriptionStyle = lipgloss.NewStyle().
		Padding(1, 2).
		Margin(1, 0).
		Border(lipgloss.RoundedBorder()).
		Width(80)
	
	// Link style
	linkStyle = lipgloss.NewStyle().
		Padding(1, 2).
		Margin(1, 0).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(linkColor).
		Foreground(linkColor)
	
	// Error style
	errorStyle = lipgloss.NewStyle().
		Bold(true).
		Padding(1, 2).
		Margin(1, 0).
		Background(clientErrorColor).
		Foreground(whiteColor).
		Border(lipgloss.RoundedBorder()).
		Align(lipgloss.Center)
	
	// List header style
	listHeaderStyle = lipgloss.NewStyle().
		Bold(true).
		Padding(1, 2).
		Margin(1, 0).
		Background(backgroundColor).
		Foreground(textColor).
		Border(lipgloss.DoubleBorder()).
		Align(lipgloss.Center).
		Width(60)
	
	// List item style
	listItemStyle = lipgloss.NewStyle().
		Padding(0, 2).
		Margin(0, 0, 0, 2)
	
	// Category header style
	categoryHeaderStyle = lipgloss.NewStyle().
		Bold(true).
		Margin(1, 0, 0, 0)
	
	// Separator style
	separatorStyle = lipgloss.NewStyle().
		Foreground(mutedColor).
		Align(lipgloss.Center).
		Margin(1, 0)
	
	// Summary style
	summaryStyle = lipgloss.NewStyle().
		Padding(0, 2).
		Margin(1, 0).
		Background(backgroundColor).
		Foreground(textColor).
		Align(lipgloss.Center)
)

// getStatusCodeColor returns the appropriate color based on HTTP status code category
func getStatusCodeColor(code int) lipgloss.Color {
	switch {
	case code >= 100 && code < 200:
		return informationalColor
	case code >= 200 && code < 300:
		return successColor
	case code >= 300 && code < 400:
		return redirectionColor
	case code >= 400 && code < 500:
		return clientErrorColor
	case code >= 500 && code < 600:
		return serverErrorColor
	default:
		return unknownColor
	}
}

// getStatusCodeCategory returns the category name for the status code
func getStatusCodeCategory(code int) string {
	switch {
	case code >= 100 && code < 200:
		return "Informational"
	case code >= 200 && code < 300:
		return "Success"
	case code >= 300 && code < 400:
		return "Redirection"
	case code >= 400 && code < 500:
		return "Client Error"
	case code >= 500 && code < 600:
		return "Server Error"
	default:
		return "Unknown"
	}
}

// displayCodeWithLipgloss displays HTTP status code information using Lipgloss styling
func displayCodeWithLipgloss(code int, info HTTPCodeInfo) {
	color := getStatusCodeColor(code)
	category := getStatusCodeCategory(code)
	
	// Display the main header with status code
	header := headerStyle.
		Foreground(color).
		BorderForeground(color).
		Render(fmt.Sprintf("HTTP %d: %s", code, info.Description))
	fmt.Println(header)
	
	// Display category badge
	badge := badgeStyle.
		Background(color).
		Render(fmt.Sprintf("📋 %s", category))
	fmt.Println(badge)
	
	// Display detailed description
	description := descriptionStyle.
		Foreground(textColor).
		BorderForeground(color).
		Render(fmt.Sprintf("📝 Description:\n\n%s", info.Detail))
	fmt.Println(description)
	
	// Display MDN link
	link := linkStyle.
		Render(fmt.Sprintf("🔗 MDN Documentation:\n%s", info.MDNLink))
	fmt.Println(link)
	
	// Add a separator
	separator := separatorStyle.
		Foreground(color).
		Render(strings.Repeat("─", 60))
	fmt.Println(separator)
}

// displayErrorWithLipgloss displays error messages using Lipgloss styling
func displayErrorWithLipgloss(message string) {
	error := errorStyle.Render(fmt.Sprintf("❌ %s", message))
	fmt.Println(error)
}

// displayListHeaderWithLipgloss displays a styled header for list commands
func displayListHeaderWithLipgloss(title string) {
	header := listHeaderStyle.Render(fmt.Sprintf("📋 %s", title))
	fmt.Println(header)
}

// displayCodeListItemWithLipgloss displays a single code item in a list
func displayCodeListItemWithLipgloss(code int, description string) {
	color := getStatusCodeColor(code)
	
	item := listItemStyle.
		Foreground(color).
		Render(fmt.Sprintf("• %d: %s", code, description))
	fmt.Println(item)
}

// displayCategoryHeaderWithLipgloss displays a category header
func displayCategoryHeaderWithLipgloss(category int, name string) {
	color := getStatusCodeColor(category * 100)
	
	header := categoryHeaderStyle.
		Foreground(color).
		Render(fmt.Sprintf("\n%dxx - %s", category, name))
	fmt.Println(header)
}

// displaySearchResultWithLipgloss displays a search result item
func displaySearchResultWithLipgloss(code int, info HTTPCodeInfo, detail string) {
	color := getStatusCodeColor(code)
	
	// Display code and description
	codeHeader := lipgloss.NewStyle().
		Bold(true).
		Foreground(color).
		Margin(1, 0, 0, 0).
		Render(fmt.Sprintf("%d: %s", code, info.Description))
	fmt.Println(codeHeader)
	
	// Display truncated detail
	detailText := lipgloss.NewStyle().
		Foreground(mutedColor).
		Margin(0, 0, 0, 2).
		Padding(0, 2).
		Render(fmt.Sprintf("📝 %s", detail))
	fmt.Println(detailText)
	
	// Display MDN link
	linkText := lipgloss.NewStyle().
		Foreground(linkColor).
		Margin(0, 0, 1, 2).
		Padding(0, 2).
		Render(fmt.Sprintf("🔗 %s", info.MDNLink))
	fmt.Println(linkText)
}

// displaySummaryWithLipgloss displays a summary message
func displaySummaryWithLipgloss(message string) {
	summary := summaryStyle.Render(message)
	fmt.Println(summary)
}
