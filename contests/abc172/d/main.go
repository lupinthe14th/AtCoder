package main

import (
	"fmt"
	"math"
)

func solution(n int) int {
	out := 0
	for i := 1; i <= n; i++ {
		out += i * count(i)
	}
	return out
}

func count(n int) int {
	out := 1
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		c := 0
		for n%i == 0 {
			c++
			n /= i
		}
		out *= (c + 1)
		if n == 1 {
			break
		}
	}
	if n != 1 {
		out *= 2
	}
	return out
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(solution(n))
}
