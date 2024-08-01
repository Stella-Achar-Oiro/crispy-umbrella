package main

import (
	"fmt"
)

func FirstWord(s string) string {
	if s == "" {
		return "\n"
	}

	f_word := ""
	for _, c := range s {
		if c == ' ' {
			break
		} else {
			f_word += string(c)
		}
	}
	return f_word + "\n"
}

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
}
