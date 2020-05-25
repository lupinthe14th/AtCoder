package main

import (
	"fmt"
)

func solution(n, d, x int, a []int) int {
	eat := 0
	for i := range a {
		for j := 0; j*a[i]+1 <= d; j++ {
			eat++
		}
	}
	return eat + x
}

func main() {
	var n, d, x int
	fmt.Scan(&n, &d, &x)
	a := make([]int, 0, n)
	for i := 0; i < n; i++ {
		t := 0
		fmt.Scan(&t)
		a = append(a, t)
	}
	fmt.Println(solution(n, d, x, a))
}
