package main

import (
	"fmt"
)

func solution(d, n int) int {
	out := 0
	switch d {
	case 0:
		if n == 100 {
			out = (n + 1)
			break
		}
		out = n
	case 1:
		if n == 100 {
			out = (n + 1) * 100
			break
		}
		out = n * 100
	case 2:
		if n == 100 {
			out = (n + 1) * 100 * 100
			break
		}
		out = n * 100 * 100
	}
	return out
}

func main() {
	var d, n int
	fmt.Scan(&d, &n)
	fmt.Println(solution(d, n))
}
