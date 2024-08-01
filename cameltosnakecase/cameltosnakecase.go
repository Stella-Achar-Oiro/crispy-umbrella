package main

import "fmt"

// CamelToSnakeCase checks if a string is camelCase and converts it to snake_case.
func CamelToSnakeCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	isCamelCase := true

	// Check if the string is valid camelCase
	for i, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			if i > 0 && s[i-1] >= 'A' && s[i-1] <= 'Z' {
				isCamelCase = false
			}
		} else if (ch < 'a' || ch > 'z') && (ch < 'A' || ch > 'Z') {
			isCamelCase = false
		}
	}

	if !isCamelCase || (s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z') {
		return s
	}

	// Convert to snake_case
	result := ""
	for i, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			if i > 0 {
				result += "_"
			}
			result += string(ch)
		} else {
			result += string(ch)
		}
	}

	return result
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))       // Output: Hello_World
	fmt.Println(CamelToSnakeCase("helloWorld"))       // Output: hello_World
	fmt.Println(CamelToSnakeCase("camelCase"))        // Output: camel_Case
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE")) // Output: CAMELtoSnackCASE
	fmt.Println(CamelToSnakeCase("camelToSnakeCase")) // Output: camel_To_Snake_Case
	fmt.Println(CamelToSnakeCase("hey2"))             // Output: hey2
}
