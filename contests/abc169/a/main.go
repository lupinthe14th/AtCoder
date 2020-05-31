package main

import (
	"fmt"
)

func solution(a, b int) int {
	return a * b
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(solution(a, b))
}
