package main

import (
	"fmt"
)

func solution(n, m int, a []int) string {
	sum := 0.0
	for _, x := range a {
		sum += float64(x)
	}
	c := 0
	for _, x := range a {
		if float64(x) >= sum/float64(4*m) {
			c++
		}
	}
	if c >= m {
		return "Yes"
	}
	return "No"
}

func main() {
	var n, m, j int
	fmt.Scan(&n, &m)
	a := make([]int, n)

	for i := range a {
		fmt.Scan(&j)
		a[i] = j
	}
	fmt.Println(solution(n, m, a))
}
