package main

import (
	"fmt"
)

// SeeAlso
// - https://atcoder.jp/contests/abc163/submissions/12126659
// - https://atcoder.jp/contests/abc163/submissions/12132835
// - https://youtu.be/MEtKSRdM4YI?t=428
func solution(n, k int) int {
	ans := 0
	for i := k; i <= n+1; i++ {
		// min := ((i - 1) * i) / 2
		// max := min + i*(n-i+1)
		// ans += max - min + 1
		ans += i*(n-i+1) + 1
		ans %= 1e9 + 7
	}
	return ans
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	fmt.Println(solution(n, k))
}
