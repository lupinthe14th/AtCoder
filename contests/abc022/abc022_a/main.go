package main

import (
	"fmt"
)

func solution(n, s, t int, a []int) int {
	x := a[0]
	c := 0
	if x >= s && t >= x {
		c++
	}
	for i := 1; i < n; i++ {
		x += a[i]
		if x >= s && t >= x {
			c++
		}
	}
	return c
}

func main() {
	var n, s, t int
	fmt.Scan(&n, &s, &t)
	a := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var w int
		fmt.Scan(&w)
		a = append(a, w)
	}
	fmt.Println(solution(n, s, t, a))
}
