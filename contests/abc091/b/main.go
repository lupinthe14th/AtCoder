package main

import (
	"fmt"
)

func solution(n, m int, s, t []string) int {
	memo := make(map[string]int)

	for _, v := range s {
		memo[v]++

	}
	for _, v := range t {
		memo[v]--
	}

	out := 0
	for _, n := range memo {
		out = max(out, n)
	}
	return out
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {
	var n int
	fmt.Scan(&n)
	s := make([]string, n)
	for i := range s {
		fmt.Scan(&s[i])
	}
	var m int
	fmt.Scan(&m)
	t := make([]string, m)
	for i := range t {
		fmt.Scan(&t[i])
	}
	fmt.Println(solution(n, m, s, t))
}
