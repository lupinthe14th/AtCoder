package main

import (
	"fmt"
)

func solution(a []int) string {
	k := a[5]
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == j {
				continue
			}
			if a[i]-a[j] > k {
				return ":("
			}
		}
	}
	return "Yay!"
}

func main() {
	a := make([]int, 0, 6)
	for i := 0; i < 6; i++ {
		t := 0
		fmt.Scan(&t)
		a = append(a, t)
	}
	fmt.Println(solution(a))
}
