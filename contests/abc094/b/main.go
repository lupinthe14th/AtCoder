package main

import (
	"fmt"
)

func solution(n, m, x int, a []int) int {

	memo := make(map[int]bool, m)
	for i := 0; i < m; i++ {
		memo[a[i]] = true
	}
	fcost := 0
	for i := x; i <= n; i++ {
		if memo[i] {
			fcost++
		}
	}

	bcost := 0
	for i := x; i >= 0; i-- {
		if memo[i] {
			bcost++
		}
	}
	return min(bcost, fcost)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	var n, m, x int
	fmt.Scan(&n, &m, &x)
	a := make([]int, 0, m)

	for i := 0; i < m; i++ {
		t := 0
		fmt.Scan(&t)
		a = append(a, t)
	}
	fmt.Println(solution(n, m, x, a))
}
