package main

import (
	"fmt"
)

func LastWord(s string) string {

	var lastword string
	wordstart := false

	for i := len(s)-1; i >= 0; i-- {
		if s[i] != ' ' {
			wordstart = true
			lastword = string(s[i]) + lastword
		} else if wordstart {
			break
		}
	}
	return lastword + "\n"
}

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}
