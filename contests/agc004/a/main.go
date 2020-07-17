package main

import (
	"fmt"
	"sort"
)

func solution(a, b, c int) int {
	if a%2 == 0 || b%2 == 0 || c%2 == 0 {
		return 0
	}
	memo := []int{a, b, c}
	sort.Ints(memo)
	return memo[0] * memo[1]
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(solution(a, b, c))
}
