package main

import (
	"fmt"
)

func solution(a, b, c, k int) int {
	if k&1 == 0 {
		return a - b
	} else {
		return b - a
	}
}

func main() {
	var a, b, c, k int
	fmt.Scan(&a, &b, &c, &k)
	fmt.Println(solution(a, b, c, k))
}
