# abort
## Instructions

Write a function that returns the median of five int arguments.
Expected function

func Abort(a, b, c, d, e int) int {

}

Usage

Here is a possible program to test your function :

package main

import (
	"fmt"
	"piscine"
)

func main() {
	middle := piscine.Abort(2, 3, 8, 5, 7)
	fmt.Println(middle)
}

And its output :

$ go run .
5
$
