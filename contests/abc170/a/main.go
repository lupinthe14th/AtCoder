package main

import (
	"fmt"
)

func solution(a []int) int {
	out := 0
	for i := range a {
		if a[i] == 0 {
			out = i + 1
		}
	}
	return out
}

func main() {
	a := make([]int, 5)
	for i := range a {
		fmt.Scan(&a[i])
	}
	fmt.Println(solution(a))
}
