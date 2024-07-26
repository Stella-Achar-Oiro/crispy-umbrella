package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]

	inter := make(map[rune]bool)
	unseen := make(map[rune]bool)

	for _, c := range str2 {
		inter[c] = true
	}

	for _, s := range str1 {
		if inter[s] && !unseen[s] {
			unseen[s] = true
			z01.PrintRune(s)
		}
	}
	z01.PrintRune('\n')
}
