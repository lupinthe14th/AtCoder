package main

import (
	"fmt"
)

func solution(n int) int {

	memo := make(map[int]int, 86)
	var rec func(n int) int

	rec = func(n int) int {
		if n == 0 {
			return 2
		}

		if n == 1 {
			return 1
		}
		out, ok := memo[n]
		if !ok {
			out = rec(n-1) + rec(n-2)
			memo[n] = out
		}
		return out
	}
	return rec(n)
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(solution(n))
}
