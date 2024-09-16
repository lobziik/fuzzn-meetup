package main

import (
	"fmt"
	"os"
)

func Reverse(input string) string {
	b := []byte(input)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	// artificialAndLowProbableCrash([]byte(input))

	return string(b)
}

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Println("You need to provide exactly one argument")
		os.Exit(1)
	}

	input := os.Args[1]
	reversed := Reverse(input)
	reversedBack := Reverse(reversed)
	fmt.Println("Original:      ", input)
	fmt.Println("Reversed:      ", reversed)
	fmt.Println("Reversed back: ", reversedBack)
}
