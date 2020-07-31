package main

import (
	"fmt"
	"sort"
)

func solution(n int, a [][2]int) int {
	out := 0

	memo := make(map[[2]int]bool)

	for i := range a {
		memo[a[i]] = true
	}

	sort.SliceStable(a, func(i, j int) bool {
		if a[i][0] == a[j][0] {
			return a[i][1] < a[j][1]
		}
		return a[i][0] < a[j][0]
	})
	var helper func(p, q [2]int)

	helper = func(p, q [2]int) {
		x := abs(p[0] - q[0])
		y := abs(p[1] - q[1])
		s := [2]int{q[0] + y, q[1] + x}
		r := [2]int{p[0] + y, p[1] + x}
		area := x*x + y*y
		if memo[s] && memo[r] {
			out = max(out, area)
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			helper(a[i], a[j])
		}
	}

	return out
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([][2]int, n)
	for i := range a {
		fmt.Scan(&a[i][0], &a[i][1])
	}
	fmt.Println(solution(n, a))
}
