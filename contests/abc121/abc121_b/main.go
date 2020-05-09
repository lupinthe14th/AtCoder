package main

import (
	"fmt"
)

func solution(n, m, c int, a [][]int, b []int) int {
	count := 0
	for i := 0; i < n; i++ {
		ans := c
		for j := 0; j < m; j++ {
			ans += a[i][j] * b[j]
		}
		if ans > 0 {
			count++
		}
	}
	return count
}

func main() {
	var n, m, c int
	fmt.Scan(&n, &m, &c)
	b := make([]int, 0, m)
	for i := 0; i < m; i++ {
		num := 0
		fmt.Scan(&num)
		b = append(b, num)
	}

	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, 0, m)
		for j := 0; j < m; j++ {
			num := 0
			fmt.Scan(&num)
			a[i] = append(a[i], num)
		}
	}
	fmt.Println(solution(n, m, c, a, b))
}
