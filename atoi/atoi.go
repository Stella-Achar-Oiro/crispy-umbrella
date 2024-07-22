package main

import (
	"fmt"
)

func Atoi(s string) int {
	sign := 1
	var number int

	for i, c := range s {
		if i == 0 && c == '-' {
			sign = -1
		} else if i == 0 && c == '+' {
			sign = 1
		} else if c >= '0' && c <= '9' {
			number = number*10 + int(c-'0')
		} else {
			return 0
		}
	}
	return sign * number
}

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}
