// Seth Solution
package main

import (
	"fmt"
)

func DigitLen(n, base int) int {
	if n < 0 {
		n = -n
	}

	if base < 2 || base > 36 {
		return -1
	}

	var count int

	for n > 0 {
		n /= base
		count++
	}
	return count
}

func main() {
	fmt.Println(DigitLen(100, 10))
	fmt.Println(DigitLen(100, 2))
	fmt.Println(DigitLen(-100, 16))
	fmt.Println(DigitLen(100, -1))
}
