package main

import (
	"fmt"
)

func solution(x int) []int {
	for a := -118; a <= 119; a++ {
		for b := -119; b <= 118; b++ {
			if a*a*a*a*a-b*b*b*b*b == x {
				return []int{a, b}
			}
		}
	}
	return []int{}
}

func main() {
	var x int
	fmt.Scan(&x)
	ans := solution(x)
	fmt.Printf("%d %d\n", ans[0], ans[1])
}
