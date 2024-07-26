package main

import (
	"fmt"
)

func HashCode(dec string) string {
	hashstring := ""

	for _, s := range dec {
		char := int(s) + len(dec) % 127
		if char >= 0 && char <= 127 {
			hashstring += string(char)
		} else {
			char += 33
		}
	}
	return hashstring
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}