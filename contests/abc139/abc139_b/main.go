package main

import (
	"fmt"
)

func solution(a, b int) int {
	if b == 1 {
		return 0
	}
	if a >= b {
		return 1
	}

	c, p := 1, a

	for p < b {
		p += a - 1
		c++
	}
	return c
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(solution(a, b))
}
