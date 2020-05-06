package main

import (
	"fmt"
)

//    展望台iから一本の道を使って辿り着けるどんな展望台よりも展望台iの方が標高が高い
// -> 展望台iから一本の道を使って辿り着ける展望台の標高の最大値よりも展望台iの標高が高い
func solution(h []int, matrix [][]int) int {
	memo := make([]int, len(h))
	for i := range matrix {
		ai := matrix[i][0] - 1
		bi := matrix[i][1] - 1
		memo[ai] = max(memo[ai], h[bi])
		memo[bi] = max(memo[bi], h[ai])
	}

	c := 0
	for i := 0; i < len(h); i++ {
		if h[i] > memo[i] {
			c++
		}
	}
	return c
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	h := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		h = append(h, num)
	}
	matrix := make([][]int, 0, m)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		matrix = append(matrix, []int{a, b})
	}
	fmt.Println(solution(h, matrix))
}
