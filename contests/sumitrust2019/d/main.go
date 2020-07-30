package main

import (
	"fmt"
)

// seen: https://atcoder.jp/contests/sumitrust2019/submissions/8717655
func solution(n int, s string) int {
	var dp [30009][4][1009]bool
	dp[0][0][0] = true

	for i := 0; i < n; i++ {
		for j := 0; j <= 3; j++ {
			for k := 0; k < 1000; k++ {
				if !dp[i][j][k] {
					continue
				}

				dp[i+1][j][k] = true
				if j <= 2 {
					dp[i+1][j+1][k*10+int(s[i]-'0')] = true
				}
			}
		}
	}

	out := 0
	for i := 0; i < 1000; i++ {
		if dp[n][3][i] {
			out++
		}
	}
	return out
}

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)
	fmt.Println(solution(n, s))
}
