package main

import (
	"fmt"
)

func solution(n int, s []int) int {
	out := 0
	for i := 0; i < 1<<uint(n); i++ {
		sum := 0
		for j := 0; j < n; j++ {
			if i>>uint(j)&1 == 1 {
				sum += s[j]
			}
		}
		if sum%10 == 0 {
			sum = 0
		}
		out = max(out, sum)
	}
	return out
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
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
