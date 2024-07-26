package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func FoldInt(f func(int, int) int, a []int, n int) {
	var result int

	tot := 0
	for _, c := range a {
		tot += c
	}

	result = f(tot, n)

	for _, c := range Itoa(result) {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func Add(a, b int) int {
	return a + b
}

func Mul(a, b int) int {
	return a * b
}

func Sub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}

func Itoa(n int) string {
	if n < 0 {
		os.Stdout.WriteString("-")
		n = -n
	} else if n == 0 {
		return "0"
	}

	digits := []rune{}

	for n > 0 {
		digit := n % 10
		digits = append([]rune{rune('0' + digit)}, digits...)
		n /= 10
	}
	return string(digits)
}

