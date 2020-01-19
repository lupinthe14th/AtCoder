package main

import (
	"fmt"
)

func solution(n, m int) string {
	if n == m {
		return "Yes"
	}
	return "No"
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	fmt.Println(solution(n, m))
}
