package main

import (
	"fmt"
	"math"
)

func solution(n, d int, x [][]int) int {
	out := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i >= j {
				continue
			}
			y := x[i]
			z := x[j]
			sum := 0
			for k := 0; k < d; k++ {
				t := y[k] - z[k]
				sum += t * t
			}
			dist := int(math.Sqrt(float64(sum)))
			if sum == dist*dist {
				out++
			}
		}
	}
	return out
}

func distance(nums []int) int {
	out := 0
	for _, n := range nums {
		out += n * n
	}
	return out
}

func main() {
	var n, d int
	fmt.Scan(&n, &d)
	x := make([][]int, n)
	for i := range x {
		x[i] = make([]int, d)
		for j := range x[i] {
			fmt.Scan(&x[i][j])
		}
	}
	fmt.Println(solution(n, d, x))
}
