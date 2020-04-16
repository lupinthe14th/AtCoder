package main

import (
	"fmt"
	"strings"
)

func rotN(n int, s string) string {
	stack := make([]string, len(s))

	for i, v := range s {
		r := v + rune(n)
		if r > 90 {
			r -= 26
		}
		stack[i] = string(r)
	}
	return strings.Join(stack, "")
}

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)
	fmt.Println(rotN(n, s))
}
