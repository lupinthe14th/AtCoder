package main

import (
	"fmt"
	"sort"
)

func solution(n int, s []int) int {
	sort.Ints(s)
	sum := 0
	for i := range s {
		sum += s[i]
	}

	if sum%10 != 0 {
		return sum
	}

	for i := range s {
		if s[i]%10 != 0 {
			return sum - s[i]
		}
	}
	return 0
}

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]int, n)
	for i := range s {
		fmt.Scan(&s[i])
	}
	fmt.Println(solution(n, s))
}
