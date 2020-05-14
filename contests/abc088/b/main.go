package main

import (
	"fmt"
	"sort"
)

func solution(n int, a []int) int {
	sort.Ints(a)

	sumA, sumB := 0, 0

	for i := 0; i < n; i++ {
		if i&1 == 0 {
			sumA += a[n-i-1]
		} else {
			sumB += a[n-i-1]
		}
	}
	return sumA - sumB
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, 0, n)
	for i := 0; i < n; i++ {
		t := 0
		fmt.Scan(&t)
		a = append(a, t)
	}
	fmt.Println(solution(n, a))
}
