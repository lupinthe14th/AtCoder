package main

import (
	"fmt"
	"math"
)

func solution(n, k int) float64 {

	memo := make([]float64, n+1)

	for i := 1; i < n+1; i++ {
		if i < k {
			x := helper(i, k)
			memo[i] = 1 / (float64(n) * math.Pow(2, x))
		} else {
			memo[i] = 1 / float64(n)
		}
	}

	out := 0.0
	for i := 1; i < n+1; i++ {
		out += memo[i]
	}
	return out
}

func helper(i, k int) float64 {
	x, out := i, 0.0
	for k > x {
		x *= 2
		out++
	}
	return out
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	fmt.Printf("%.12f\n", solution(n, k))
}
