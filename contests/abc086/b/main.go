package main

import (
	"fmt"
	"math"
)

func solution(a, b int) string {
	d := helper(a, b) + b

	n := math.Sqrt(float64(d))
	if n == math.Floor(n) {
		return "Yes"
	}
	return "No"
}

func helper(a, b int) int {
	e := 1
	for b > 0 {
		b /= 10
		e *= 10
	}
	return a * e
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(solution(a, b))
}
