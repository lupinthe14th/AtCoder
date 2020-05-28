package main

import (
	"fmt"
)

func solution(a, b, c, d int) (x3, y3, x4, y4 int) {
	x := c - a
	y := d - b
	return c - y, d + x, a - y, b + x
}

func main() {
	var x1, y1, x2, y2 int
	fmt.Scan(&x1, &y1, &x2, &y2)
	fmt.Println(solution(x1, y1, x2, y2))
}
