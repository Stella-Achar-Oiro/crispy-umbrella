package main 

import (
	"fmt"
	"strconv"
)

func ZipString(s string) string {
	if len(s) == 0 {
		return ""
	}

	q := ""
	count := 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			q += strconv.Itoa(count) + string(s[i-1])
			count = 1
		}
	}

	q += strconv.Itoa(count) + string(s[len(s)-1])

	return q
}

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}