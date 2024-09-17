package main

import (
	"testing"
	"unicode/utf8"
)

func FuzzReversed(f *testing.F) {
	// initial corpus
	testcases := []string{
		"Hello, world",
		" ",
		"",
		"!12345",
		"longlonglonglonglonglonglonglonginput",
	}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, input string) {
		rev, err := Reverse(input)
		if err != nil {
			return
		}

		revBack, err := Reverse(rev)
		if err != nil {
			return
		}

		// Prop 1
		// String should be equal to original one after it reversed back
		if input != revBack {
			t.Errorf("Before: %q, after: %q", input, revBack)
		}

		// Prop 2
		// String should remain a valid UTF-8 string after reverse, if original input was a valid UTF-8 string
		if utf8.ValidString(input) && !utf8.ValidString(rev) {
			t.Errorf("Invalid UTF-8 string after reverseing %q: %q", input, rev)
		}
	})
}
