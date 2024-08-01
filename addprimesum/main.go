package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	num, err := Atoi(os.Args[1])
	if !err {
		return
	}

	psum := 0
	for i := 2; i <= num; i++ {
		if IsPrime(i) {
			psum += i
		}
	}
	for _, c := range Itoa(psum) {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Atoi(s string) (int, bool) {
	sign := 1
	var numb int

	for idx, chr := range s {
		if idx == 0 && chr == '-' {
			sign = -1
		} else if idx == 0 && chr == '+' {
			sign = 1
		} else if chr >= '0' && chr <= '9' {
			numb = numb*10 + int(chr-'0')
		} else {
			return 0, false
		}
	}
	return sign * numb, true
}

func Itoa(n int) string {
	if n < 0 {
		os.Stdout.WriteString("-")
		n = -n
	} else if n == 0 {
		return "0"
	}

	digits := []rune{}

	for n > 0 {
		digit := n % 10
		digits = append([]rune{rune('0' + digit)}, digits...)
		n /= 10
	}
	return string(digits)
}