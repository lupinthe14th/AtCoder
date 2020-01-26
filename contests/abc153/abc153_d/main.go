package main

import (
	"fmt"
)

// refarence: https://atcoder.jp/contests/abc153/submissions/9741217
func solution(h int) int {
	if h == 0 {
		return 0
	}
	if h == 1 {
		return 1
	}
	return 1 + 2*solution(h/2)
}

func main() {
	var h int
	fmt.Scan(&h)
	fmt.Println(solution(h))
}
