package main

import (
	"fmt"
)

func solution(r, g, b, n int) int {
	c := 0
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			v := r*i + g*j
			if n >= v && (n-v)%b == 0 {
				c++
			}
		}
	}
	return c
}

func main() {
	var r, g, b, n int
	fmt.Scan(&r, &g, &b, &n)
	fmt.Println(solution(r, g, b, n))
}
