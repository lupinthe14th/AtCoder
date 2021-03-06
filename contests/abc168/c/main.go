package main

import (
	"fmt"
	"math"
)

func solution(a, b, h, m float64) float64 {
	M := 6 * m
	H := (h / 12 * 360) + (m / 60 * 30)
	A := abs(H, M)
	if A >= 180 {
		A = 360 - A
	}
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2) - 2*a*b*math.Cos(A*(math.Pi/180)))
}

func main() {
	var a, b, h, m int
	fmt.Scan(&a, &b, &h, &m)
	fmt.Println(solution(float64(a), float64(b), float64(h), float64(m)))
}

func abs(x, y float64) float64 {
	if x > y {
		return x - y
	}
	return y - x
}
