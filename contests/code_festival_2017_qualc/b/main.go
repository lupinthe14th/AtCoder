package main

import (
	"fmt"
)

func solution(n int, a []int) int {
	all, bad := 1, 1
	for i := 0; i < n; i++ {
		all *= 3
		if a[i]%2 == 0 {
			bad *= 2
		}
	}
	return all - bad
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	fmt.Println(solution(n, a))
}
