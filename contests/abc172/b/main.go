package main

import (
	"fmt"
)

func solution(s, t string) int {
	out := 0
	for i := range s {
		if s[i] != t[i] {
			out++
		}
	}
	return out
}

func main() {
	var s, t string
	fmt.Scan(&s, &t)
	fmt.Println(solution(s, t))
}
