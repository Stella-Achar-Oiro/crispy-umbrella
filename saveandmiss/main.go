package main

import (
	"fmt"
)

func SaveAndMiss(arg string, num int) string {
	newStr := ""
	count := 0

	for i := 0; i <= len(arg) -1; i++ {
		count++
		newStr += string(arg[i])
		if count == num {
			count = 0
			i += num
		}
	}
	return newStr
}

func main() {
	fmt.Println(SaveAndMiss("123456789", 3))
	fmt.Println(SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Println(SaveAndMiss("", 3))
	fmt.Println(SaveAndMiss("hello you all ! ", 0))
	fmt.Println(SaveAndMiss("what is your name?", 0))
	fmt.Println(SaveAndMiss("go Exercise Save and Miss", -5))
}