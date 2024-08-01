package main

import "fmt"

func AtoiBase(s string, base string) int {
	if !IsValidBase(base) {
		return 0
	}

	baseLen := len(base)
	baseMap := make(map[rune]int)
	for i, r := range base {
		baseMap[r] = i
	}

	q := 0
	for _, v := range s {
		value, ok := baseMap[v]
		if !ok {
			return 0
		}
		q = q * baseLen + value
	}
	return q
}

func IsValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	seen := make(map[rune]bool)
	for _, r := range base {
		if r == '+' || r == '-' || seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}
