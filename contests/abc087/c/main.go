package main

import (
	"fmt"
)

func solution(n int, a [2][]int) int {
	out := -1 << 31
	for k := n - 1; k >= 0; k-- {
		i := 0
		sum := 0
		for j := 0; j < n; j++ {
			sum += a[i][j]
			if k == j {
				i++
				sum += a[i][j]
			}
		}
		out = max(out, sum)
	}
	return out
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var n int
	fmt.Scan(&n)
	a := [2][]int{}
	for i := range a {
		a[i] = make([]int, n)
		for j := range a[i] {
			fmt.Scan(&a[i][j])
		}
	}
	fmt.Println(solution(n, a))
}
