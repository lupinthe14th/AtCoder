package main

import (
	"fmt"
)

func solution(n, k int) int {
	t := n % k
	return min(t, k-t)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	fmt.Println(solution(n, k))
}
