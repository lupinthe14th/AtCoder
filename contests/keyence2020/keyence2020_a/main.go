package main

import (
	"fmt"
)

func solution(h, w, n int) int {
	max := max(h, w)
	var ans int = 1
	for max*ans < n {
		ans++
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var h, w, n int
	fmt.Scan(&h, &w, &n)
	fmt.Println(solution(h, w, n))
}
