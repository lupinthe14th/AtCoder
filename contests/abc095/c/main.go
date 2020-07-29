package main

import (
	"fmt"
)

func solution(a, b, c, x, y int) int {
	d := a*x + b*y
	s := 0
	if x < y {
		s = 2*c*x + ((y - x) * b)
	} else {
		s = 2*c*y + ((x - y) * a)
	}
	p := c * 2 * max(x, y)

	return min(p, min(d, s))
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
