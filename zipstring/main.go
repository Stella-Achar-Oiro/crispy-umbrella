package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func ZipString(s string) string {
	if len(s) == 0 {
		return ""
	}

	result := ""

	var currentChar rune
	count := 0

	for i, char := range s {
		if i == 0 {
			currentChar = char
			count = 1
		} else {
			if char == currentChar {
				count++
			} else {
				result += Itoa(count) + string(currentChar)
				currentChar = char
				count = 1
			}
		}
	}
	result += Itoa(count) + string(currentChar)
	return result
}

func Itoa(n int) string {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	} else if n == 0 {
		return "0"
	}

	digits := []rune{}

	for n > 0 {
		digit := n % 10
		digits = append([]rune{rune('0' + digit)}, digits...)
		n /= 10
	}
	return string(digits)
}

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}
