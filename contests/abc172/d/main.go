package main

import (
	"fmt"
)

func solution(n int) int {
	out := 0
	for i := 1; i <= n; i++ {
		out += i * g(n/i)
	}
	return out
}

func g(n int) int {
	return n * (n + 1) / 2
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(solution(n))
}
