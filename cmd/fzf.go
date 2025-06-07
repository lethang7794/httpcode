package cmd

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
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

type httpCodeItem struct {
	code     int
	desc     string
	category string
}

func runFzfSearch(withPreview bool) {
	// Prepare data for fuzzy search
	var items []httpCodeItem
	for code, desc := range httpCodes {
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
		items = append(items, httpCodeItem{code, desc, category})
	}

	// Sort items by code
	sort.Slice(items, func(i, j int) bool {
		return items[i].code < items[j].code
	})

	// Initialize terminal screen
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Printf("Error creating screen: %v\n", err)
		return
	}
	if err := screen.Init(); err != nil {
		fmt.Printf("Error initializing screen: %v\n", err)
		return
	}
	defer screen.Fini()

	// Get screen dimensions
	width, height := screen.Size()

	// Main loop variables
	pattern := ""
	selected := 0
	offset := 0
	maxItems := height - 4 // Reserve space for header and prompt
	quit := false

	for !quit {
		screen.Clear()

		// Draw header
		headerText := "HTTP Status Codes - Press Ctrl+C to exit, Enter to select"
		for i, r := range headerText {
			if i >= width {
				break
			}
			screen.SetContent(i, 0, r, nil, tcell.StyleDefault.Foreground(tcell.ColorGreen).Bold(true))
		}

		// Draw search prompt
		promptText := "Search: " + pattern
		for i, r := range promptText {
			if i >= width {
				break
			}
			screen.SetContent(i, 1, r, nil, tcell.StyleDefault)
		}

		// Filter items based on pattern
		var filtered []httpCodeItem
		if pattern == "" {
			filtered = items
		} else {
			for _, item := range items {
				searchText := fmt.Sprintf("%d %s %s", item.code, item.desc, item.category)
				if fuzzyMatch(searchText, pattern) {
					filtered = append(filtered, item)
				}
			}
		}

		// Draw items
		for i := 0; i < maxItems && i+offset < len(filtered); i++ {
			item := filtered[i+offset]
			var text string
			if withPreview {
				text = fmt.Sprintf("%d: %s", item.code, item.desc)
			} else {
				text = fmt.Sprintf("%d: %s", item.code, item.desc)
			}
			
			style := tcell.StyleDefault
			
			// Highlight selected item
			if i+offset == selected {
				style = style.Background(tcell.ColorBlue).Foreground(tcell.ColorWhite)
			}
			
			// Draw item
			for j, r := range text {
				if j >= width/2 - 1 {
					break
				}
				screen.SetContent(j, i+3, r, nil, style)
			}
		}

		// Draw preview if enabled
		if withPreview && len(filtered) > 0 && selected < len(filtered) {
			item := filtered[selected]
			
			previewX := width / 2
			previewWidth := width - previewX - 1
			
			// Draw preview border
			for i := 0; i < height; i++ {
				screen.SetContent(previewX-1, i, '│', nil, tcell.StyleDefault)
			}
			
			// Draw preview content
			previewLines := []string{
				"HTTP Status Code: " + strconv.Itoa(item.code),
				"",
				"Description: " + item.desc,
				"",
				"Category: " + item.category + " (" + strconv.Itoa(item.code/100) + "xx)",
			}
			
			for i, line := range previewLines {
				if i >= height {
					break
				}
				
				style := tcell.StyleDefault
				if i == 0 || i == 2 || i == 4 {
					style = style.Foreground(tcell.ColorGreen).Bold(true)
				}
				
				for j, r := range line {
					if j >= previewWidth {
						break
					}
					screen.SetContent(previewX+j, i+1, r, nil, style)
				}
			}
		}

		screen.Show()

		// Handle input
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyCtrlC:
				quit = true
			case tcell.KeyEnter:
				if len(filtered) > 0 && selected < len(filtered) {
					item := filtered[selected]
					screen.Fini()
					fmt.Printf("%d: %s\n", item.code, item.desc)
					return
				}
			case tcell.KeyUp:
				if selected > 0 {
					selected--
					if selected < offset {
						offset = selected
					}
				}
			case tcell.KeyDown:
				if selected < len(filtered)-1 {
					selected++
					if selected >= offset+maxItems {
						offset = selected - maxItems + 1
					}
				}
			case tcell.KeyHome:
				selected = 0
				offset = 0
			case tcell.KeyEnd:
				if len(filtered) > 0 {
					selected = len(filtered) - 1
					if selected >= maxItems {
						offset = selected - maxItems + 1
					} else {
						offset = 0
					}
				}
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				if len(pattern) > 0 {
					pattern = pattern[:len(pattern)-1]
					selected = 0
					offset = 0
				}
			case tcell.KeyRune:
				pattern += string(ev.Rune())
				selected = 0
				offset = 0
			}
		case *tcell.EventResize:
			width, height = screen.Size()
			maxItems = height - 4
		}
	}
}

// Simple fuzzy matching algorithm
func fuzzyMatch(str, pattern string) bool {
	str = strings.ToLower(str)
	pattern = strings.ToLower(pattern)
	
	if pattern == "" {
		return true
	}
	
	if str == "" {
		return false
	}
	
	// Check if all characters in pattern appear in str in the same order
	strIdx := 0
	for _, patChar := range pattern {
		found := false
		for strIdx < len(str) {
			if rune(str[strIdx]) == patChar {
				found = true
				strIdx++
				break
			}
			strIdx++
		}
		if !found {
			return false
		}
	}
	
	return true
}

func init() {
	rootCmd.AddCommand(fzfCmd)
	rootCmd.AddCommand(fzfSearchCmd)
}
