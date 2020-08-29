package main

import (
	"fmt"
)

func solution(s, t string) int {
	out := 1 << 31
	count := func(a, b string) int {
		c := 0
		for i := range a {
			if a[i] != b[i] {
				c++
			}
		}
		return c
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	for i := 0; i < len(s)-len(t)+1; i++ {
		out = min(out, count(s[i:len(t)+i], t))
	}

	return out
}

func main() {
	var s, t string
	fmt.Scan(&s, &t)
	fmt.Println(solution(s, t))
}
