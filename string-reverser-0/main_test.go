package main

import "testing"

func TestReverse(t *testing.T) {
	testcases := []struct {
		input, expected string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"", ""},
		{"!12345", "54321!"},
		{"longlonglonglonglonglonglonglonginput", "tupnignolgnolgnolgnolgnolgnolgnolgnol"},
	}
	for _, tc := range testcases {
		rev := Reverse(tc.input)
		if rev != tc.expected {
			t.Errorf("Reverse: %q, want %q", rev, tc.expected)
		}
	}
}
