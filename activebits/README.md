# activebits
## Instructions
Write a function, ActiveBits, that returns the number of active bits (bits with the value 1) in the binary representation of an integer number.

Expected function
func ActiveBits(n int) uint {

}
Usage
Here is a possible program to test your function :

package main

import (
	"fmt"
	piscine ".."
)

func main() {
	nbits := piscine.ActiveBits(7)
	fmt.Println(nbits)
}
And its output :

student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
3
student@ubuntu:~/[[ROOT]]/test$