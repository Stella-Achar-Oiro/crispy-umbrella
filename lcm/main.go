package main

import (
	"fmt"
)

// Gcd calculates the greatest common divisor of two numbers.
func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a % b
	}
	return a
}

// Lcm calculates the least common multiple of two numbers.
func Lcm(first, second int) int {
	if first == 0 || second == 0 {
		return 0
	}
	gcd := Gcd(first, second)
	return (first * second) / gcd
}

func main() {
	fmt.Println(Lcm(2, 7)) // Output: 14
	fmt.Println(Lcm(0, 4)) // Output: 0
	fmt.Println(Lcm(6, 8)) // Output: 24
}
