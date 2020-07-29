package main

import (
	"fmt"
)

func solution(n, m int, a [][]int) int {

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	out := 0
	for i := 0; i < m; i++ {
		for j := i + 1; j < m; j++ {
			sum := 0
			for k := 0; k < n; k++ {
				sum += max(a[k][i], a[k][j])
			}
			out = max(out, sum)
		}
	}
	return out
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, m)
		for j := range a[i] {
			fmt.Scan(&a[i][j])
		}
	}
	fmt.Println(solution(n, m, a))
}
