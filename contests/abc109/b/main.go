package main

import (
	"fmt"
)

func solution(n int, w []string) string {
	memo := make(map[string]bool, n)
	memo[w[0]] = true

	pre := w[0][len(w[0])-1]
	for i := 1; i < n; i++ {
		if _, ok := memo[w[i]]; ok || pre != w[i][0] {
			return "No"
		}
		pre = w[i][len(w[i])-1]
		memo[w[i]] = true
	}
	return "Yes"
}

func main() {
	var n int
	fmt.Scan(&n)
	w := make([]string, n)
	for i := range w {
		fmt.Scan(&w[i])
	}
	fmt.Println(solution(n, w))
}
