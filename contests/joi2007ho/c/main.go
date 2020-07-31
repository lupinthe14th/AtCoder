package main

import (
	"fmt"
)

func solution(n int, a [][2]int) int {
	out := 0

	memo := make(map[[2]int]int)

	for i := range a {
		memo[a[i]] = i
	}
	var helper func(p, q [2]int)

	helper = func(p, q [2]int) {
		x := abs(p[0] - q[0])
		y := abs(p[1] - q[1])
		s1 := [2]int{q[0] + y, q[1] + x}
		s2 := [2]int{q[0] - y, q[1] - x}
		r1 := [2]int{p[0] + y, p[1] + x}
		r2 := [2]int{p[0] - y, p[1] - x}
		area := x*x + y*y
		if _, ok := memo[s1]; ok {
			if _, ok := memo[r1]; ok {
				out = max(out, area)
			}
		} else if _, ok := memo[s2]; ok {
			if _, ok := memo[r2]; ok {
				out = max(out, area)
			}
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
