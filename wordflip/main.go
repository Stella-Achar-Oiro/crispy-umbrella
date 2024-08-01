package main

import (
	"fmt"
)

func WordFlip(str string) string {
	if str == "" {
		return "Invalid Output\n"
	} 

	nwSlice := splitWhiteSpaces(str)
	nwWord := []string{}
	word := ""

	for i := len(nwSlice)-1; i >= 0; i-- {
		nwWord = append(nwWord, nwSlice[i]) 
	}

	for i, c := range nwWord {
		if i < len(nwWord)-1 {
			word += c + " "
		} else {
			word += c
		}
	}
	return word + "\n"
}


func splitWhiteSpaces(s string) []string {
	wordSlice := []string{}

	var start int
	for i, c := range s {
		if c == ' ' {
			if start < i {
				wordSlice = append(wordSlice, s[start:i])
			}
			start = i + 1
		}
	}
	if start < len(s) {
		wordSlice = append(wordSlice, s[start:])
	}
	return wordSlice
}

func main() {
	fmt.Print(WordFlip("First second last"))
	fmt.Print(WordFlip(""))
	fmt.Print(WordFlip("     "))
	fmt.Print(WordFlip(" hello  all  of  you! "))
}
