package main

import (
	"fmt"
)

func solution(n int) int {
	return n/2 + n%2
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(solution(n))
}
