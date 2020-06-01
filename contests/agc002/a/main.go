package main

import (
	"fmt"
)

func solution(a, b int) string {
	out := ""
	switch {
	case a > 0:
		out = "Positive"
	case a < 0 && b > 0:
		out = "Zero"
	case b < 0:
		if abs(-a-b)%2 != 0 {
			out = "Positive"
		} else {
			out = "Negative"
		}
	}
	return out
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(solution(a, b))
}
