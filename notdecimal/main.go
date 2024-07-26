package main

import (
	"fmt"
)

func NotDecimal(dec string) string {
	if dec == "" {
		return "\n"
	}

	decimal := ""

	for _, c := range dec {
		if c == '.' {
			continue
		} else {
			decimal += string(c)
		}
	}

	// fmt.Println(decimal)

	for i, n := range decimal {
		if n >= '1' && n <= '9' && decimal[0] == '0' {
			decimal = decimal[i:]
		}
	}

	if !isNumber(dec) {
		return dec + "\n"
	}
	return decimal + "\n"
}

func isNumber(s string) bool {
	for _, c := range s {
		if c >= '0' && c <= '9' {
			return true
		} else {
			return false
		}
	}
	return false
}

func main() {
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.01255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
}
