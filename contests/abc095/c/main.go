package main

import (
	"fmt"
)

func solution(a, b, c, x, y int) int {
	out := max(x, y) * c * 2

	out = min(out, a*x+b*y)

	tmp := c * 2 * min(x, y)
	t := min(x, y)
	x -= t
	y -= t

	tmp += a*x + b*y

	return min(out, tmp)
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
