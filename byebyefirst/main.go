package main

import (
	"fmt"
)

func ByeByeFirst(strings []string) []string {
	if len(strings) <= 1 {
		return nil
	}
	return strings[1:]
}

func main() {
	fmt.Println(ByeByeFirst([]string{}))
	fmt.Println(ByeByeFirst([]string{"one arg"}))
	fmt.Println(ByeByeFirst([]string{"first", "second"}))
	fmt.Println(ByeByeFirst([]string{"", "abcd", "efg"}))
}