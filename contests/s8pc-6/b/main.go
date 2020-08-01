package main

import (
	"fmt"
)

func solution(n int, a, b []int) int64 {
	var out int64 = 1<<63 - 1
	min := func(x, y int64) int64 {
		if x < y {
			return x
		}
		return y
	}

	abs := func(x int) int64 {
		x64 := int64(x)
		if x64 < 0 {
			return -x64
		}
		return x64
	}
	solve := func(s, t int) int64 {
		var out int64
		for i := 0; i < n; i++ {
			out += abs(s - a[i])
			out += abs(a[i] - b[i])
			out += abs(b[i] - t)
		}
		return out
	}

	for i := range a {
		for j := range b {
			out = min(out, solve(a[i], b[j]))
		}
	}

	return out
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i], &b[i])
	}
	fmt.Println(solution(n, a, b))
}
