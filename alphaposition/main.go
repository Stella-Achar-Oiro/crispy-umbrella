package main

import "fmt"

func AlphaPosition(c rune) int {
    pos := 0

	if c >= 'a' && c <= 'z' {
		for i := 'a'; i <= 'z'; i++ {
			pos++
			if i == c {
				break
			}
		}
	} else if c >= 'A' && c <= 'Z' {
		for i := 'A'; i <= 'Z'; i++ {
			pos++
			if i == c {
				break
			}
		}
	} else {
		return -1
	}
	return pos
}

func main(){
    fmt.Println(AlphaPosition('a'))
    fmt.Println(AlphaPosition('z'))
    fmt.Println(AlphaPosition('B'))
    fmt.Println(AlphaPosition('Z'))
    fmt.Println(AlphaPosition('0'))
    fmt.Println(AlphaPosition(' '))
}