package main

import (
	"fmt"
	"sort"
)

func solution(n, m int, matrix [][2]int) int {
	sort.SliceStable(matrix, func(i, j int) bool {
		if matrix[i][0] == matrix[j][0] {
			return matrix[i][1] > matrix[j][1]
		}
		return matrix[i][0] < matrix[j][0]
	})

	out := 0
	for i := range matrix {
		if m >= matrix[i][1] {
			m -= matrix[i][1]
			out += matrix[i][0] * matrix[i][1]
		} else {
			out += matrix[i][0] * m
			m = 0
		}
		if m <= 0 {
			break
		}
	}
	return out
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	matrix := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&matrix[i][0], &matrix[i][1])
	}
	fmt.Println(solution(n, m, matrix))
}
