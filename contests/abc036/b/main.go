package main

import (
	"fmt"
	"strings"
)

func solution(n int, s [][]string) [][]string {
	ans := make([][]string, n)
	for i := range ans {
		ans[i] = make([]string, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ans[j][n-1-i] = s[i][j]
		}
	}
	return ans
}

func main() {
	var n int
	fmt.Scan(&n)

	s := make([][]string, n)
	for i := range s {
		var t string
		fmt.Scan(&t)
		s[i] = make([]string, n)
		for j := range t {
			s[i][j] = string(t[j])
		}
	}
	for _, text := range solution(n, s) {
		fmt.Println(strings.Join(text, ""))
	}
}
