package main

import (
	"fmt"
)

func palindromePhilia(s string) int {
	var c int
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			c++
		}
	}
	return c
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(palindromePhilia(s))
}
