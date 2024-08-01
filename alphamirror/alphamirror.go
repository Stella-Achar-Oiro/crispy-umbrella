package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]
	q := make([]rune, len(args))

	for i, v := range args {
		if v >= 'a' && v <= 'z' {
			q[i] = 'z' - (v - 'a')
		} else if v >= 'A' && v <= 'Z' {
			q[i] = 'Z' - (v - 'A')
		} else {
			q[i] = v
		}
	}
	PrintStr(string(q))
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
