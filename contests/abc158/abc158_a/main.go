package main

import (
	"fmt"
)

func solution(s string) string {
	for i := 1; i < 3; i++ {
		if s[i-1] != s[i] {
			return "Yes"
		}
	}
	return "No"
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(solution(s))
}
