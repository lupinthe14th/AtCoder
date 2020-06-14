package main

import (
	"fmt"
)

func solution(s string) string {
	return s[:3]
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(solution(s))
}
