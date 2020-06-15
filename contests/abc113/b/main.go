package main

import (
	"fmt"
)

func solution(n, t, a int, h []int) int {
	lo := 1 << 31
	out := 0
	for i := range h {
		ave := 1000*t - (h[i] * 6)
		diff := abs(1000*a - ave)
		if lo > diff {
			out = i
			lo = diff
		}

	}
	return out + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var n, t, a int
	fmt.Scan(&n, &t, &a)
	h := make([]int, n)
	for i := range h {
		fmt.Scan(&h[i])
	}
	fmt.Println(solution(n, t, a, h))
}
