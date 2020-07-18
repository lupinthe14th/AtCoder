package main

import (
	"fmt"
	"math"
)

func solution(n, k int) int {
	return k * int(math.Pow(float64(k-1), float64(n-1)))
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	fmt.Println(solution(n, k))
}
