package main

import (
	"fmt"
)

func solution(a, b, c, x, y int) int {
	out := 1 << 31

	for i := 0; i <= 1e5; i++ {
		out = min(out, i*2*c+max(0, x-i)*a+max(0, y-i)*b)
	}
	return out
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var a, b, c, x, y int
	fmt.Scan(&a, &b, &c, &x, &y)
	fmt.Println(solution(a, b, c, x, y))
}
