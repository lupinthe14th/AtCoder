package main

import (
	"fmt"
)

func solution(n int, a []int) int {
	var out int = 1

	for _, num := range a {
		if num == 0 {
			return 0
		}
	}
	for _, num := range a {
		if num <= 1e18/out {
			out *= num
		} else {
			return -1
		}
	}
	return out
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var t int
		fmt.Scan(&t)
		a = append(a, t)
	}
	fmt.Println(solution(n, a))
}
