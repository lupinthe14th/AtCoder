package main

import (
	"fmt"
)

func solution(n int) int {
	s := 0
	for s < n {
		s += 1000
		if s == n {
			break
		}
	}
	return s - n
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(solution(n))
}
