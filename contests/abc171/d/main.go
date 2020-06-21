package main

import (
	"fmt"
)

func solution(n int, a []int, q int, b, c []int) []uint64 {
	out := make([]uint64, q)
	for i := 0; i < q; i++ {
		var sum uint64
		for j := range a {
			if a[j] == b[i] {
				a[j] = c[i]
			}
			sum += uint64(a[j])
		}
		out[i] = sum
	}
	return out
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	var q int
	fmt.Scan(&q)
	b := make([]int, q)
	c := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Scan(&b[i], &c[i])
	}
	for _, t := range solution(n, a, q, b, c) {
		fmt.Println(t)
	}
}
