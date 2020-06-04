package main

import (
	"fmt"
)

func solution(n int) int {
	var L [87]int
	L[0] = 2
	L[1] = 1
	for i := 2; i <= n; i++ {
		L[i] = L[i-1] + L[i-2]
	}
	return L[n]
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(solution(n))
}
