package main

import (
	"fmt"
)

func solution(n int, h []int) string {
	for i := n - 1; i > 0; i-- {
		if h[i] >= h[i-1] {
			continue
		}
		h[i-1]--
		if h[i] < h[i-1] {
			return "No"
		}
	}
	return "Yes"
}

func main() {
	var n int
	fmt.Scan(&n)
	h := make([]int, n)
	for i := range h {
		fmt.Scan(&h[i])
	}
	fmt.Println(solution(n, h))
}
