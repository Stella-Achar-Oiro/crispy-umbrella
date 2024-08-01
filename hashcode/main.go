package main

import (
	"fmt"
)

// HashCode computes the hashed string based on the given hash equation.
func HashCode(dec string) string {
	size := len(dec)
	hashedString := ""

	for _, char := range dec {
		// Compute ASCII value + size of the string
		newValue := int(char) + size
		
		// Apply modulo 127 to keep the result within ASCII range
		newValue = newValue % 127
		
		// Ensure the result is printable
		if newValue < 33 {
			newValue += 33
		}
		
		// Append the resulting character to the hashed string
		hashedString += string(rune(newValue))
	}

	return hashedString
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
	fmt.Println(HashCode("14 Avril 1912"))
}
