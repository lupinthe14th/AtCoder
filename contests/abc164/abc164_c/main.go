package main

import (
	"fmt"
)

func solution(n int, ss []string) int {
	gifts := make(map[string]int)
	for _, s := range ss {
		gifts[s]++
	}
	return len(gifts)
}

func main() {
	var n int
	fmt.Scan(&n)
	ss := make([]string, 0, n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		ss = append(ss, s)
	}
	fmt.Println(solution(n, ss))
}
