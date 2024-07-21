package main

import (
	"fmt"
)

func ActiveBits(n int) uint {
	count := uint(0)
	for n != 0 {
		count += uint(n & 1)
		n >>= 1
	}
	return count
}

func main() {
	nbits := ActiveBits(7)
	fmt.Println(nbits)
}
