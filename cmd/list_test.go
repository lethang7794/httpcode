package cmd

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestListCommand(t *testing.T) {
	tests := []struct {
		name         string
		args         []string
		wantContains []string
		wantError    bool
	}{
		{
			name: "list all codes",
			args: []string{},
			wantContains: []string{
				"All HTTP Status Codes",
				"1xx - Informational",
				"2xx - Success",
				"3xx - Redirection",
				"4xx - Client Error",
				"5xx - Server Error",
				"200: OK",
				"404: Not Found",
			},
			wantError: false,
		},
		{
			name: "list 4xx codes",
			args: []string{"4xx"},
			wantContains: []string{
				"4xx - Client Error",
				"400: Bad Request",
				"404: Not Found",
			},
			wantError: false,
		},
		{
			name: "list 2xx codes",
			args: []string{"2xx"},
			wantContains: []string{
				"2xx - Success",
				"200: OK",
				"201: Created",
			},
			wantError: false,
		},
		{
			name: "invalid category format",
			args: []string{"4x"},
			wantContains: []string{
				"Invalid category",
				"Use 1xx, 2xx, 3xx, 4xx, or 5xx",
			},
			wantError: true,
		},
		{
			name: "invalid category number",
			args: []string{"6xx"},
			wantContains: []string{
				"Invalid category",
				"Use 1xx, 2xx, 3xx, 4xx, or 5xx",
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new list command for each test
			cmd := &cobra.Command{
				Use: "list [category]",
				Run: func(cmd *cobra.Command, args []string) {
					listCmd.Run(cmd, args)
				},
			}
			
			cmd.SetArgs(tt.args)
			
			stdout, _ := captureOutput(func() {
				cmd.Execute()
			})
			
			for _, want := range tt.wantContains {
				if !strings.Contains(stdout, want) {
					t.Errorf("Expected '%s' in output, got: %s", want, stdout)
				}
			}
		})
	}
}

func TestListCodes(t *testing.T) {
	tests := []struct {
		name         string
		category     string
		wantContains []string
		wantError    bool
	}{
		{
			name:     "empty category lists all",
			category: "",
			wantContains: []string{
				"All HTTP Status Codes",
				"1xx - Informational",
				"2xx - Success",
			},
			wantError: false,
		},
		{
			name:     "4xx category",
			category: "4xx",
			wantContains: []string{
				"4xx - Client Error",
				"400:",
				"404:",
			},
			wantError: false,
		},
		{
			name:     "case insensitive",
			category: "4XX",
			wantContains: []string{
				"4xx - Client Error",
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout, _ := captureOutput(func() {
				listCodes(tt.category)
			})

			for _, want := range tt.wantContains {
				if !strings.Contains(stdout, want) {
					t.Errorf("Expected '%s' in output, got: %s", want, stdout)
				}
			}
		})
	}
}
