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

	tmp := make([][]int, 0, l)
	for key, val := range seen {
		tmp = append(tmp, []int{key, val})
	}
	sort.SliceStable(tmp, func(i, j int) bool {
		return tmp[i][1] < tmp[j][1]
	})
	out := 0
	for i := range tmp[:l-k] {
		out += tmp[i][1]
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
