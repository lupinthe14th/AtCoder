package main

import (
	"fmt"
)

func solution(n, k int, x []int) int {
	sum := 0
	for i := 0; i < n; i++ {
		a := x[i] * 2
		b := (k - x[i]) * 2
		sum += min(a, b)
	}
	return sum
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	x := make([]int, n)

	for i := range x {
		t := 0
		fmt.Scan(&t)
		x[i] = t
	}
	fmt.Println(solution(n, k, x))
}
