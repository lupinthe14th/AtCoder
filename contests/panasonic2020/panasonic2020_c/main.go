package main

import (
	"fmt"
)

func solution(a, b, c int64) string {
	d := c - a - b
	if d > 0 && d*d > 4*a*b {
		return "Yes"
	}
	return "No"
}

func main() {
	var a, b, c int64
	fmt.Scan(&a, &b, &c)
	fmt.Println(solution(a, b, c))
}
