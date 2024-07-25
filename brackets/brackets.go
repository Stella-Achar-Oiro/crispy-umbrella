package main

import (
	"fmt"
	"os"
)

func Brackets(s string) string {
	q := []rune{}
	bracketPairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, r := range s {
		switch r {
		case '(', '[', '{':
			q = append(q, r)
		case ')', ']', '}':
			if len(q) == 0 || q[len(q)-1] != bracketPairs[r] {
				return "Error"
			}
			q = q[:len(q)-1]
		}
	}
	if len(q) == 0 {
		return "OK"
	}
	return "Error"
}

func main() {
	if len(os.Args) <= 1 {
		return
	}
	for _, arg := range os.Args[1:] {
		fmt.Println(Brackets(arg))
	}
}
