package main

import (
	"fmt"
)

func solution(n int, s string) int {
	out := -1 << 31
	for i := 1; i < n; i++ {
		x := count(s[:i])
		y := count(s[i:])
		t := 0
		for j := 0; j < 26; j++ {
			if x[j] > 0 && y[j] > 0 {
				t++
			}
		}
		out = max(out, t)
	}
	return out
}

func count(s string) [26]int {
	var a [26]int
	for _, r := range s {
		a[r-'a']++
	}
	return a
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)
	fmt.Println(solution(n, s))
}
