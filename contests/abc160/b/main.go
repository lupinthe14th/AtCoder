package main

import (
	"fmt"
)

func solution(x int) int {
	var c, ans int
	c = x / 500
	ans = c * 1000
	x = x - 500*c
	c = x / 5
	return ans + c*5
}

func main() {
	var x int
	fmt.Scan(&x)
	fmt.Println(solution(x))
}
