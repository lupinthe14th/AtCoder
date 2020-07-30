package main

import (
	"fmt"
)

// seen: https://atcoder.jp/contests/sumitrust2019/submissions/8695967
func solution(n int, s string) int {
	out := 0

	for i := 0; i < 1000; i++ {
		b := [3]byte{byte(i / 100), byte((i / 10) % 10), byte(i % 10)}
		f := 0
		for j := 0; j < n; j++ {
			if s[j] == ('0' + b[f]) {
				f++
			}
			if f == 3 {
				break
			}
		}
		if f == 3 {
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
