package main

import (
	"fmt"
)

func WeAreUnique(str1, str2 string) int {
	if str1 == "" || str2 == "" {
		return -1
	}
	
	unique := make(map[rune]bool)
	unseen := make(map[rune]bool)
	count := 0

	for _, c := range str1 {
		unique[c] = true
	}

	for _, c := range str2 {
		unseen[c] = true
	}

	for _, c := range str1 {
		if !unseen[c] {
			count++
		}
	}
	for _, c := range str2 {
		if !unique[c] {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}
