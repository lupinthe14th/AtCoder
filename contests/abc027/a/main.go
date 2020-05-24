package main

import (
	"fmt"
	"sort"
)

func solution(n, x int, a []int) int {
	sort.Ints(a)
	for i, c := range a {
		if x < c {
			return i
		}
		x -= c
	}
	if x == 0 {
		return n
	}
	return n - 1
}

func main() {
	var n, x int
	fmt.Scan(&n, &x)
	a := make([]int, 0, n)
	for i := 0; i < n; i++ {
		t := 0
		fmt.Scan(&t)
		a = append(a, t)
	}
	fmt.Println(solution(n, x, a))
}
