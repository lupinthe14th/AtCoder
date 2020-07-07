package main

import (
	"fmt"
)

const MOD = 1e9 + 7

func solution(n int) int {
	ans := 1
	for i := 1; i <= n; i++ {
		ans = ans * i % MOD
	}
	return ans
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(solution(n))
}
