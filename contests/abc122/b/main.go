package main

import (
	"fmt"
)

func solution(s string) int {
	l := len(s)
	out := 0
	for i := 0; i <= l; i++ {
		for j := i + 1; j <= l; j++ {
			t := s[i:j]
			if isACGT(t) {
				out = max(out, len(t))
			}
		}
	}
	return out
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func isACGT(s string) bool {
	for _, r := range s {
		if r == 'A' || r == 'C' || r == 'G' || r == 'T' {
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(solution(s))
}
