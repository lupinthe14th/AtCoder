package main

import (
	"fmt"
)

func solution(l int) float64 {
	var x float64 = float64(l) / 3
	return x * x * x
}

func main() {
	var l int
	fmt.Scan(&l)
	fmt.Println(solution(l))
}
