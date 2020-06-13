package main

import (
	"fmt"
)

func solution(n, m int, ab, cd [][]int) []int {
	out := make([]int, n)

	for i := range ab {
		near := 1 << 31
		for j := range cd {
			cur := abs(ab[i][0]-cd[j][0]) + abs(ab[i][1]-cd[j][1])
			if near > cur {
				near = cur
				out[i] = j + 1
			}
		}
	}
	return out
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	ab := make([][]int, n)
	for i := range ab {
		ab[i] = make([]int, 2)
		fmt.Scan(&ab[i][0], &ab[i][1])
	}
	cd := make([][]int, m)
	for i := range cd {
		cd[i] = make([]int, 2)
		fmt.Scan(&cd[i][0], &cd[i][1])
	}
	for _, num := range solution(n, m, ab, cd) {
		fmt.Println(num)
	}
}
