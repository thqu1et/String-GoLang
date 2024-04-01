package stringutil

import (
	"String/stringutil/stringutil"
	"testing"
)

func TestSymbolCount(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"hello", 5},
		{"world", 5},
		{"", 0},
		{"こんにちは", 5}, // Example with Unicode characters
	}

	for _, test := range tests {
		result := stringutil.SymbolCount(test.input)
		if result != test.expected {
			t.Errorf("SymbolCount(%s) = %d; want %d", test.input, result, test.expected)
		}
	}
}
