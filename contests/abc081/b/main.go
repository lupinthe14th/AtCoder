package main

import (
	"fmt"
)

func solution(n int, a []int) int {
	out := 0
	for {
		flag := true
		for i := range a {
			if a[i]%2 != 0 {
				flag = false
				break
			}
			a[i] >>= 1
		}
		if !flag {
			break
		}
		out++
	}
	return out
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, 0, n)
	for i := 0; i < n; i++ {
		t := 0
		fmt.Scan(&t)
		a = append(a, t)
	}
	fmt.Println(solution(n, a))
}
