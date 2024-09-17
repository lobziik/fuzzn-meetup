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
		rev, err := Reverse(tc.input)
		if err != nil {
			t.Errorf("Reverse returned error! input: %q err: %q", tc.input, err)
		}
		if rev != tc.expected {
			t.Errorf("Reverse: %q, want %q", rev, tc.expected)
		}
	}
}
