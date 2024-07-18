package main

import "testing"

func TestRev(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"", ""},
		{"a", "a"},
		{"hello world", "dlrow olleh"},
	}

	for _, tc := range testCases {
		result := Rev(tc.input)
		if result != tc.expected {
			t.Errorf("Rev(%q) : %q; expected : %q", tc.input, result, tc.expected)
		}
	}
}
