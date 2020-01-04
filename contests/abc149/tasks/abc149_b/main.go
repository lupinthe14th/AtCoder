package main

import (
	"fmt"
)

func solution(a, b, k int64) []int64 {
	ans := make([]int64, 0, 2)
	if a >= k {
		a -= k
		ans = []int64{a, b}
	} else if a < k {
		k -= a
		a = 0
		b -= k
		if b < 0 {
			b = 0
		}
		ans = []int64{a, b}
	}
	return ans
}

func main() {
	var a, b, k int64
	fmt.Scan(&a, &b, &k)
	ans := solution(a, b, k)
	fmt.Printf("%d %d", ans[0], ans[1])
}
