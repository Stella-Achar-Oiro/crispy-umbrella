package main

import (
	"unicode"

	"github.com/01-edu/z01"
)

func PrintMemory(arr [10]byte) {
	for i, b := range arr {
		printhex(b)
		if (i + 1)%4 == 0 || i == len(arr)-1 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}
	for _, b := range arr {
		if unicode.IsGraphic(rune(b)) {
			z01.PrintRune(rune(b))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}

func printhex(b byte) {
	printDigits := "0123456789abcdef"
	z01.PrintRune(rune(printDigits[b>>4]))
	z01.PrintRune(rune(printDigits[b&0x0f]))
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}