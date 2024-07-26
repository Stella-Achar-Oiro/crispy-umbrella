package main

import (
	"fmt"
)

func CanJump(num []uint) bool {
	if len(num) == 0 {
		return false
	}

	if len(num) == 1 {
		return true
	}

	position := 0

	for position < len(num)-1 {
		steps := int(num[position])
		if steps == 0 || position + steps >= len(num) {
			return false
		}
		position += steps
		}
	
	return position == len(num)-1
}

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2))

	input3 := []uint{0}
	fmt.Println(CanJump(input3))
}