package main

import (
	"fmt"
	"sort"
)

func solution(a, b, c int) int {
	memo := []int{a, b, c}
	sort.Ints(memo)
	hi := memo[2]
	lo := memo[0]
	mi := memo[1]
	if (hi-(hi-mi+lo))%2 == 1 {
		return hi - mi + 2 + (hi-(hi-mi+lo))/2

	}
	return hi - mi + (hi-(hi-mi+lo))/2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(solution(a, b, c))
}
