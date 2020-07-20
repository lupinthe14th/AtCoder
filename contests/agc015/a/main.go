package main

import (
	"fmt"
)

func solution(n, a, b int) int {
	if a > b {
		return 0
	}

	if a == b {
		return 1
	}
	l := n - 2
	if l < 0 {
		return 0
	}
	return l*b - l*a + 1
}

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	fmt.Println(solution(n, a, b))
}
