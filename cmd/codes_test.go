package cmd

import (
	"fmt"
	"strings"
	"testing"
)

func TestHTTPCodesInfo(t *testing.T) {
	// Test that we have some basic HTTP codes
	expectedCodes := []int{200, 201, 301, 302, 400, 401, 403, 404, 500, 502, 503}
	
	for _, code := range expectedCodes {
		t.Run(fmt.Sprintf("code_%d_exists", code), func(t *testing.T) {
			info, exists := httpCodesInfo[code]
			if !exists {
				t.Errorf("HTTP code %d should exist in httpCodesInfo", code)
				return
			}
			
			if info.Description == "" {
				t.Errorf("HTTP code %d should have a description", code)
			}
			
			if info.Detail == "" {
				t.Errorf("HTTP code %d should have detail information", code)
			}
			
			if info.MDNLink == "" {
				t.Errorf("HTTP code %d should have an MDN link", code)
			}
			
			// Check that MDN link is properly formatted
			if !strings.HasPrefix(info.MDNLink, "https://developer.mozilla.org/") {
				t.Errorf("HTTP code %d MDN link should start with https://developer.mozilla.org/, got: %s", code, info.MDNLink)
			}
		})
	}
}

func TestHTTPCodesInfoStructure(t *testing.T) {
	// Test that all codes have the required fields
	for code, info := range httpCodesInfo {
		t.Run(fmt.Sprintf("code_%d_structure", code), func(t *testing.T) {
			if info.Description == "" {
				t.Errorf("HTTP code %d missing Description", code)
			}
			
			if info.Detail == "" {
				t.Errorf("HTTP code %d missing Detail", code)
			}
			
			if info.MDNLink == "" {
				t.Errorf("HTTP code %d missing MDNLink", code)
			}
			
			// Validate code is in correct range
			if code < 100 || code >= 600 {
				t.Errorf("HTTP code %d is outside valid range (100-599)", code)
			}
		})
	}
}

func TestHTTPCodeCategories(t *testing.T) {
	// Test that we have codes in each major category
	categories := map[int]string{
		1: "1xx",
		2: "2xx", 
		3: "3xx",
		4: "4xx",
		5: "5xx",
	}
	
	categoryCounts := make(map[int]int)
	
	for code := range httpCodesInfo {
		category := code / 100
		categoryCounts[category]++
	}
	
	for category, name := range categories {
		t.Run(fmt.Sprintf("category_%s_exists", name), func(t *testing.T) {
			count := categoryCounts[category]
			if count == 0 {
				t.Errorf("No HTTP codes found in category %s", name)
			} else {
				t.Logf("Found %d codes in category %s", count, name)
			}
		})
	}
}

func TestSpecificHTTPCodes(t *testing.T) {
	tests := []struct {
		code        int
		description string
		category    string
	}{
		{200, "OK", "Success"},
		{404, "Not Found", "Client Error"},
		{500, "Internal Server Error", "Server Error"},
		{301, "Moved Permanently", "Redirection"},
		{100, "Continue", "Informational"},
	}
	
	for _, tt := range tests {
		t.Run(fmt.Sprintf("code_%d", tt.code), func(t *testing.T) {
			info, exists := httpCodesInfo[tt.code]
			if !exists {
				t.Errorf("HTTP code %d should exist", tt.code)
				return
			}
			
			if info.Description != tt.description {
				t.Errorf("HTTP code %d description = %q, want %q", tt.code, info.Description, tt.description)
			}
			
			// Test category function
			category := getStatusCodeCategory(tt.code)
			if category != tt.category {
				t.Errorf("HTTP code %d category = %q, want %q", tt.code, category, tt.category)
			}
		})
	}
}
