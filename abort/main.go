package main

import (
	"fmt"
)

func Abort(a, b, c, d, e int) int {
	var n = []int{a, b, c, d, e}

	for i := 0; i < len(n)-1; i ++ {
		for j := i + 1; j < len(n); j++ {
			if n[i] > n[j] {
				n[i], n[j] = n[j], n[i]
			}
		}
	}
	median := len(n)/2
	return n[median]
}

func main() {
	middle := Abort(2, 3, 8, 5, 7)
	fmt.Println(middle)
}