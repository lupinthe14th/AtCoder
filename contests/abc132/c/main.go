package main

import (
	"fmt"
	"sort"
)

func solution(n int, d []int) int {
	sort.Ints(d)
	l := len(d)
	return d[l/2] - d[l/2-1]
}

func main() {
	var n int
	fmt.Scan(&n)
	d := make([]int, n)
	for i := 0; i < n; i++ {
		var t int
		fmt.Scan(&t)
		d[i] = t
	}
	fmt.Println(solution(n, d))
}
