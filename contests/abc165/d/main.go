package main

import (
	"fmt"
)

func solution(a, b, n int) int {
	x := n

	if n >= b-1 {
		x = b - 1
	}

	return a * (x % b) / b
}

func main() {
	var a, b, n int
	fmt.Scan(&a, &b, &n)
	fmt.Println(solution(a, b, n))
}
