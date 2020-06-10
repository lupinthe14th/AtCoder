package main

import (
	"fmt"
)

func solution(n, m int, ab [][]int) []int {
	out := make([]int, n)

	for i := range ab {
		out[ab[i][0]-1]++
		out[ab[i][1]-1]++
	}
	return out
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	ab := make([][]int, m)
	for i := range ab {
		ab[i] = make([]int, 2)
		fmt.Scan(&ab[i][0], &ab[i][1])
	}
	for _, v := range solution(n, m, ab) {
		fmt.Println(v)
	}
}
