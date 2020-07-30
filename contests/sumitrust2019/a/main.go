package main

import (
	"fmt"
)

func solution(m1, d1, m2, d2 int) int {
	if m1+1 == m2 {
		return 1
	}
	return 0
}

func main() {
	var m1, d1, m2, d2 int
	fmt.Scan(&m1, &d1, &m2, &d2)
	fmt.Println(solution(m1, d1, m2, d2))
}
