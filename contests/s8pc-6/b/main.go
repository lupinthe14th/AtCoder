package main

import (
	"fmt"
	"sort"
)

func solution(n int, ab [][2]int) int64 {
	in := make([]int, n)
	out := make([]int, n)
	for i := range ab {
		in[i] = ab[i][0]
		out[i] = ab[i][1]
	}
	sort.Ints(in)
	sort.Ints(out)
	gateIn := int64(in[(n-1)/2])
	gateOut := int64(out[(n-1)/2])

	var ans int64
	for i := range ab {
		var sum, a, b int64
		a = int64(ab[i][0])
		b = int64(ab[i][1])
		switch {
		case gateIn <= a && b <= gateOut:
			sum = gateOut - gateIn
		case a < gateIn && b <= gateOut:
			sum = 2*(gateIn-a) + gateOut - gateIn
		case gateIn <= a && gateOut < b:
			sum = gateOut - gateIn + 2*(b-gateOut)
		case a < gateIn && gateOut < b:
			sum = gateOut - gateIn + 2*(gateIn-a) + 2*(b-gateOut)
		case a < gateIn && b < gateIn:
			sum = 2*(gateIn-a) + gateOut - gateIn
		case gateOut < a && gateOut < b:
			sum = gateOut - gateIn + 2*(b-gateOut)
		}
		ans += sum
	}
	return ans
}

func main() {
	var n int
	fmt.Scan(&n)
	ab := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ab[i][0], &ab[i][1])
	}
	fmt.Println(solution(n, ab))
}
