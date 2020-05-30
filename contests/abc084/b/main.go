package main

import (
	"fmt"
)

func solution(a, b int, s string) string {
	if s[a] != '-' {
		return "No"
	}
	for i := 0; i < a+b+1; i++ {
		if i == a {
			continue
		}
		n := s[i] - '0'
		if 0 > n || n > 9 {
			return "No"
		}
	}
	return "Yes"
}

func main() {
	var a, b int
	var s string
	fmt.Scan(&a, &b, &s)
	fmt.Println(solution(a, b, s))
}
