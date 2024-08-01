package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] == "" {
		z01.PrintRune('\n')
		return
	}

	str := splitWhiteSpaces(os.Args[1])
	for _, c := range str {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func splitWhiteSpaces(s string) string {
	var newSlice []string
	var start int
	newStr := ""

	for i, c := range s {
		if c == ' ' {
			if start < i {
				newSlice = append(newSlice, s[start:i])
			}
			start = i + 1
		}
	}
	if start < len(s) {
		newSlice = append(newSlice, s[start:])
	}
	for i, c := range newSlice {
		if i < len(newSlice)-1 {
			newStr += c + " "
		} else {
			newStr += c
		}
	}
	return newStr
}