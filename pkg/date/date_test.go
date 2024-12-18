package date

import (
	"testing"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		logEntry    string
		shouldMatch bool
		shouldError bool
	}{
		// Test valid formats
		{"Valid YYYY format", "2024", "[17/06/24, 12:34:56] Some message", true, false},
		{"Valid YYYY-MM format", "2024-06", "[17/06/24, 12:34:56] Some message", true, false},
		{"Valid YYYY-MM-DD format", "2024-06-17", "[17/06/24, 12:34:56] Some message", true, false},

		// Test invalid formats
		{"Invalid format: only year prefix", "20", "", false, true},
		{"Invalid format: YYYY-MM extra character", "2024-06-XX", "", false, true},
		{"Invalid format: random string", "invalid-date", "", false, true},
		{"Invalid format: incomplete date", "2024-06-", "", false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			regex, err := Filter(tt.input)
			if tt.shouldError {
				if err == nil {
					t.Errorf("Expected an error for input %q but got nil", tt.input)
				}
				return
			}
			if err != nil {
				t.Errorf("Unexpected error for input %q: %v", tt.input, err)
				return
			}

			matches := regex.MatchString(tt.logEntry)
			if matches != tt.shouldMatch {
				t.Errorf("Regex mismatch for input %q. Expected %v, got %v", tt.input, tt.shouldMatch, matches)
			}
		})
	}
}
