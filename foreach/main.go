package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func ForEach(f func(int), a []int) {
	for _, c := range a {
		f(c)
	}
	z01.PrintRune('\n')
}

func PrintNbr(n int) {
	fmt.Print(n)
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	ForEach(PrintNbr, a)
}