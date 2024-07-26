package main 

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]

	i := 0
	j := 0

	str3 := ""

	for i < len(str1) && j < len(str2) {
		if str1[i] == str2[j] {
			str3 += string(str1[i])
			i++
		}
		j++
	}
	if str1 == str3 {
		z01.PrintRune('1')
	} else {
		z01.PrintRune('0')
	}
	z01.PrintRune('\n')
}