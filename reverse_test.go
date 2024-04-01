package stringutil_test

import (
	"String/stringutil/stringutil"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		input, expected string
	}{
		{"hello", "olleh"},
		{"world", "dlrow"},
		{"", ""},
	}

	for _, test := range tests {
		result := stringutil.Reverse(test.input)
		if result != test.expected {
			t.Errorf("Reverse(%s) = %s; want %s", test.input, result, test.expected)
		}
	}
}
