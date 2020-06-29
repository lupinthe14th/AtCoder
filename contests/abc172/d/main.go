package main

import (
	"fmt"
)

func solution(n int) int {
	memo := make([]int, n+1)

	for i := 1; i <= n; i++ {
		for j := 1; i*j <= n; j++ {
			memo[i*j]++
		}
	}

	out := 0
	for i := 1; i <= n; i++ {
		out += i * memo[i]
	}
	return out
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(solution(n))
}
