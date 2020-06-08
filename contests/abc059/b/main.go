package main

import (
	"fmt"
)

func solution(a, b string) string {
	if len(a) > len(b) {
		return "GREATER"
	}
	if len(a) < len(b) {
		return "LESS"
	}
	if a > b {
		return "GREATER"
	}
	if a == b {
		return "EQUAL"
	}
	return "LESS"
}

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	fmt.Println(solution(a, b))
}
