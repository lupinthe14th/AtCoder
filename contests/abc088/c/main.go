package main

import (
	"fmt"
)

func solution(c [3][3]int) string {
	var a, b [3]int
	a[0] = 0
	for i := 0; i < 3; i++ {
		b[i] = c[0][i]
	}
	for i := 1; i < 3; i++ {
		a[i] = c[i][0] - b[0]
	}

	for i := range c {
		for j := range c[i] {
			if c[i][j] != a[i]+b[j] {
				return "No"
			}
		}
	}
	return "Yes"
}

func main() {
	var c [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scan(&c[i][j])
		}
	}
	fmt.Println(solution(c))
}
