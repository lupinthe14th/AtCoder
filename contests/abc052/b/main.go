package main

import (
	"fmt"
)

func solution(n int, s string) int {
	out, x := 0, 0
	for i := 0; i < n; i++ {
		if s[i] == 'I' {
			x++
		} else {
			x--
		}
		out = max(out, x)
	}
	return out
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)
	fmt.Println(solution(n, s))
}
