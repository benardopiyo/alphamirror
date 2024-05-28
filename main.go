package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1]

	var result string
	for _, char := range args {
		if char >= 'a' && char <= 'z' {
			char = ('z' - (char - 'a'))
			result = result + string(char)
		} else if char >= 'A' && char <= 'Z' {
			char = ('Z' - (char - 'A'))
			result = result + string(char)
		} else {
			result = result + string(char)
		}
	}

	for _, val := range result {
		z01.PrintRune(val)
	}
	z01.PrintRune(10)
}
