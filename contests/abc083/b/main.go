package main

import (
	"fmt"
)

func solution(n, a, b int) int {
	out := 0
	for i := 1; i <= n; i++ {
		t := sumDigit(i)
		if a <= t && t <= b {
			out += i
		}
	}
	return out
}

func sumDigit(n int) int {
	out := 0
	for n > 0 {
		out += n % 10
		n /= 10
	}
	return out
}

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	fmt.Println(solution(n, a, b))
}
