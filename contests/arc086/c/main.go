package main

import (
	"fmt"
	"sort"
)

func solution(n, k int, a []int) int {
	seen := make(map[int]int, 0)

	for i := range a {
		seen[a[i]]++
	}

	l := len(seen)
	if l <= k {
		return 0
	}

	tmp := make([]int, 0, l)
	for _, v := range seen {
		tmp = append(tmp, v)
	}
	sort.Ints(tmp)
	out := 0
	for i := 0; i < l-k; i++ {
		out += tmp[i]
	}
	return out
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	fmt.Println(solution(n, k, a))
}
