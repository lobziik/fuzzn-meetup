package main

import (
	"errors"
	"unicode/utf8"
)

func goodReverse(input string) (string, error) {
	if !utf8.ValidString(input) {
		return input, errors.New("input is not valid UTF-8")
	}
	b := []rune(input)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b), nil
}
