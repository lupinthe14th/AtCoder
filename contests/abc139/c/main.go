package main

import (
	"fmt"
)

func solution(n int, h []int) int {
	memo := make([]int, n)
	for i := 1; i < n; i++ {
		if h[i-1] >= h[i] {
			memo[i] = memo[i-1] + 1
		}
	}
	out := -1 << 31
	for i := range memo {
		out = max(out, memo[i])
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
	fmt.Scan(&n)
	h := make([]int, n)
	for i := 0; i < n; i++ {
		a := 0
		fmt.Scan(&a)
		h[i] = a
	}
	fmt.Println(solution(n, h))
}
