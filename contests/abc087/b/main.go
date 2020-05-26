package main

import (
	"fmt"
)

func solution(a, b, c, x int) int {
	out := 0
	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for k := 0; k <= c; k++ {
				if x == 500*i+100*j+50*k {
					out++
				}
			}
		}
	}
	return out
}

func main() {
	var a, b, c, x int
	fmt.Scan(&a, &b, &c, &x)
	fmt.Println(solution(a, b, c, x))
}
