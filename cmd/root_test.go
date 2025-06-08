package cmd

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

// captureOutput captures stdout and stderr for testing
func captureOutput(f func()) (string, string) {
	oldStdout := os.Stdout
	oldStderr := os.Stderr
	
	rOut, wOut, _ := os.Pipe()
	rErr, wErr, _ := os.Pipe()
	
	os.Stdout = wOut
	os.Stderr = wErr
	
	outC := make(chan string)
	errC := make(chan string)
	
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, rOut)
		outC <- buf.String()
	}()
	
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, rErr)
		errC <- buf.String()
	}()
	
	f()
	
	wOut.Close()
	wErr.Close()
	os.Stdout = oldStdout
	os.Stderr = oldStderr
	
	stdout := <-outC
	stderr := <-errC
	
	return stdout, stderr
}

func TestRootCommand(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantCode bool
		wantHelp bool
		wantErr  bool
	}{
		{
			name:     "valid status code",
			args:     []string{"404"},
			wantCode: true,
			wantHelp: false,
			wantErr:  false,
		},
		{
			name:     "invalid status code",
			args:     []string{"999"},
			wantCode: false,
			wantHelp: false,
			wantErr:  true,
		},
		{
			name:     "non-numeric argument",
			args:     []string{"invalid"},
			wantCode: false,
			wantHelp: true,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new root command for each test
			cmd := &cobra.Command{
				Use: "httpcode [code]",
				Run: func(cmd *cobra.Command, args []string) {
					rootCmd.Run(cmd, args)
				},
			}
			
			cmd.SetArgs(tt.args)
			
			stdout, stderr := captureOutput(func() {
				cmd.Execute()
			})
			
			if tt.wantCode {
				if !strings.Contains(stdout, "HTTP "+tt.args[0]) {
					t.Errorf("Expected HTTP status code %s in output, got: %s", tt.args[0], stdout)
				}
			}
			
			if tt.wantHelp {
				if !strings.Contains(stdout, "Unknown command") {
					t.Errorf("Expected help message in output, got: %s", stdout)
				}
			}
			
			if tt.wantErr {
				if !strings.Contains(stdout, "not found") {
					t.Errorf("Expected error message for invalid code, got: %s", stdout)
				}
			}
			
			// Check stderr is empty when not expecting errors
			if !tt.wantErr && stderr != "" {
				t.Errorf("Unexpected stderr output: %s", stderr)
			}
		})
	}
}

func TestLookupCode(t *testing.T) {
	tests := []struct {
		name         string
		code         int
		wantContains []string
		wantError    bool
	}{
		{
			name: "valid 404 code",
			code: 404,
			wantContains: []string{
				"HTTP 404:",
				"Not Found",
				"Category:",
				"Client Error",
				"Description:",
				"Documentation:",
			},
			wantError: false,
		},
		{
			name: "valid 200 code",
			code: 200,
			wantContains: []string{
				"HTTP 200:",
				"OK",
				"Category:",
				"Success",
			},
			wantError: false,
		},
		{
			name: "invalid code",
			code: 999,
			wantContains: []string{
				"HTTP status code 999 not found",
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout, _ := captureOutput(func() {
				lookupCode(tt.code)
			})

			for _, want := range tt.wantContains {
				if !strings.Contains(stdout, want) {
					t.Errorf("Expected '%s' in output, got: %s", want, stdout)
				}
			}
		})
	}
}
