package main

import (
	"errors"
	"fmt"
	"os"
	"unicode/utf8"
)

func Reverse(input string) (string, error) {
	if !utf8.ValidString(input) {
		return input, errors.New("input is not valid UTF-8")
	}
	b := []rune(input)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b), nil
}

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Println("You need to provide exactly one argument")
		os.Exit(1)
	}

	input := os.Args[1]
	reversed, err := Reverse(input)
	if err != nil {
		fmt.Println("Error happened during the reverse: ", err.Error())
	}
	reversedBack, err := Reverse(reversed)
	if err != nil {
		fmt.Println("Error happened during the reversing back: ", err.Error())
	}
	fmt.Println("Original:      ", input)
	fmt.Println("Reversed:      ", reversed)
	fmt.Println("Reversed back: ", reversedBack)
}
