package main

import (
	"fmt"
	"sort"
)

func solution(q, h, s, d, n int) int {
	//  25: q
	//  50: h
	// 100: s
	// 200: d
	// 100: n
	n = 100 * n
	out := 0
	a := [][]int{{25, q, 8 * q}, {50, h, 4 * h}, {100, s, 2 * s}, {200, d, d}}
	sort.SliceStable(a, func(i, j int) bool {
		if a[i][2] == a[j][2] {
			return a[i][0] < a[j][0]
		}
		return a[i][2] < a[j][2]
	})

	for i := range a {
		if n >= a[i][0] {
			out += a[i][1] * (n / a[i][0])
			n %= a[i][0]
		}
	}
	return out
}

func main() {
	var q, h, s, d, n int
	fmt.Scan(&q, &h, &s, &d, &n)
	fmt.Println(solution(q, h, s, d, n))
}
