package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	input, words, currentWord := os.Args[1], []string{}, []rune{}

	for _, char := range input {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			currentWord = append(currentWord, char)
		} else {
			if len(currentWord) > 0 {
				words = append(words, string(currentWord))
				currentWord = []rune{}
			}
		}
	}
	if len(currentWord) > 0 {
		words = append(words, string(currentWord))
	}

	if len(words) == 0 {
		return
	}

	for i, word := range words {
		os.Stdout.WriteString(word)
		if i < len(words)-1 {
			os.Stdout.WriteString("   ")
		}
	}
	z01.PrintRune('\n')
}
