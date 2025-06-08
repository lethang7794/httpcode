package cmd

import (
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
)

func TestGetStatusCodeColor(t *testing.T) {
	tests := []struct {
		name     string
		code     int
		expected lipgloss.Color
	}{
		{
			name:     "1xx informational",
			code:     100,
			expected: informationalColor,
		},
		{
			name:     "2xx success",
			code:     200,
			expected: successColor,
		},
		{
			name:     "3xx redirection",
			code:     301,
			expected: redirectionColor,
		},
		{
			name:     "4xx client error",
			code:     404,
			expected: clientErrorColor,
		},
		{
			name:     "5xx server error",
			code:     500,
			expected: serverErrorColor,
		},
		{
			name:     "unknown code",
			code:     999,
			expected: unknownColor,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getStatusCodeColor(tt.code)
			if result != tt.expected {
				t.Errorf("getStatusCodeColor(%d) = %v, want %v", tt.code, result, tt.expected)
			}
		})
	}
}

func TestDisplayCodeWithLipgloss(t *testing.T) {
	tests := []struct {
		name         string
		code         int
		info         HTTPCodeInfo
		wantContains []string
	}{
		{
			name: "404 not found",
			code: 404,
			info: HTTPCodeInfo{
				Description: "Not Found",
				Detail:      "The server cannot find the requested resource.",
				MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404",
			},
			wantContains: []string{
				"HTTP 404",
				"Not Found",
				"Class:",
				"Client Error",
				"Description:",
				"Docs:",
			},
		},
		{
			name: "200 ok",
			code: 200,
			info: HTTPCodeInfo{
				Description: "OK",
				Detail:      "The request has succeeded.",
				MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/200",
			},
			wantContains: []string{
				"HTTP 200",
				"OK",
				"Class:",
				"Success",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout, _ := captureOutput(func() {
				displayCodeWithLipgloss(tt.code, tt.info)
			})

			for _, want := range tt.wantContains {
				if !strings.Contains(stdout, want) {
					t.Errorf("Expected '%s' in output, got: %s", want, stdout)
				}
			}
		})
	}
}

func TestDisplayErrorWithLipgloss(t *testing.T) {
	tests := []struct {
		name         string
		message      string
		wantContains []string
	}{
		{
			name:    "error message",
			message: "Test error message",
			wantContains: []string{
				"❌",
				"Test error message",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout, _ := captureOutput(func() {
				displayErrorWithLipgloss(tt.message)
			})

			for _, want := range tt.wantContains {
				if !strings.Contains(stdout, want) {
					t.Errorf("Expected '%s' in output, got: %s", want, stdout)
				}
			}
		})
	}
}

func TestDisplayListHeaderWithLipgloss(t *testing.T) {
	tests := []struct {
		name         string
		title        string
		wantContains []string
	}{
		{
			name:  "list header",
			title: "Test Header",
			wantContains: []string{
				"📋",
				"Test Header",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout, _ := captureOutput(func() {
				displayListHeaderWithLipgloss(tt.title)
			})

			for _, want := range tt.wantContains {
				if !strings.Contains(stdout, want) {
					t.Errorf("Expected '%s' in output, got: %s", want, stdout)
				}
			}
		})
	}
}

func TestDisplayCodeListItemWithLipgloss(t *testing.T) {
	tests := []struct {
		name         string
		code         int
		description  string
		wantContains []string
	}{
		{
			name:        "list item",
			code:        404,
			description: "Not Found",
			wantContains: []string{
				"404:",
				"Not Found",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout, _ := captureOutput(func() {
				displayCodeListItemWithLipgloss(tt.code, tt.description)
			})

			for _, want := range tt.wantContains {
				if !strings.Contains(stdout, want) {
					t.Errorf("Expected '%s' in output, got: %s", want, stdout)
				}
			}
		})
	}
}

func TestDisplayCategoryHeaderWithLipgloss(t *testing.T) {
	tests := []struct {
		name         string
		category     int
		categoryName string
		wantContains []string
	}{
		{
			name:         "4xx category",
			category:     4,
			categoryName: "Client Error",
			wantContains: []string{
				"4xx",
				"Client Error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout, _ := captureOutput(func() {
				displayCategoryHeaderWithLipgloss(tt.category, tt.categoryName)
			})

			for _, want := range tt.wantContains {
				if !strings.Contains(stdout, want) {
					t.Errorf("Expected '%s' in output, got: %s", want, stdout)
				}
			}
		})
	}
}

func TestDisplaySummaryWithLipgloss(t *testing.T) {
	tests := []struct {
		name         string
		message      string
		wantContains []string
	}{
		{
			name:    "summary message",
			message: "Found 5 results",
			wantContains: []string{
				"Found 5 results",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout, _ := captureOutput(func() {
				displaySummaryWithLipgloss(tt.message)
			})

			for _, want := range tt.wantContains {
				if !strings.Contains(stdout, want) {
					t.Errorf("Expected '%s' in output, got: %s", want, stdout)
				}
			}
		})
	}
}
