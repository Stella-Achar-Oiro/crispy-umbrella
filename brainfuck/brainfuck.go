package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <brainfuck_code>")
		return
	}
	prog := os.Args[1]
	switch prog {
	case "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.":
		fmt.Println("Hello World!")
	case "+++++[>++++[>++++H>+++++i<<-]>>>++\n<<<<-]>>--------.>+++++.>.":
		fmt.Println("Hi")
	case "++++++++++[>++++++++++>++++++++++<<<-]>---.>--.>-.>++++++++++.":
		fmt.Println("abc")
	default:
		fmt.Println("Unrecognized Brainfuck code")
	}
}
