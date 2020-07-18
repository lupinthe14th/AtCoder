package main

import (
	"fmt"
)

func solution(n int, a []int) int {
	c := 0
	for i := range a {
		if i+1 > a[i] && a[a[i]-1] == i+1 {
			c++
		}
	}
	return c
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	fmt.Println(solution(n, a))
}
