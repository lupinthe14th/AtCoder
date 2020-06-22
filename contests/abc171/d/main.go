package main

import (
	"fmt"
)

func solution(n int, a []int, q int, b, c []int) []uint64 {
	var memo [1e5 + 1]int
	var sum uint64
	for i := range a {
		memo[a[i]]++
		sum += uint64(a[i])
	}
	out := make([]uint64, q)
	for i := 0; i < q; i++ {
		out[i] = sum
		num := memo[b[i]]
		if num > 0 {
			sum += uint64(num * (c[i] - b[i]))
			out[i] = sum
			memo[b[i]] = 0
			memo[c[i]] += num
		}
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
