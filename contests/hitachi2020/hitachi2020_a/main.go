package main

import (
	"fmt"
)

func solution(s string) string {
	l := len(s)

	if l < 2 {
		return "No"
	}
	if l%2 == 1 {
		return "No"
	}
	for i := 2; i <= l; i += 2 {
		if s[i-2:i] != "hi" {
			return "No"
		}
	}
	return "Yes"
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(solution(s))
}
