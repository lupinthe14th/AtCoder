package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func solution(n int, p [][2]int) float64 {
	out := 0.0
	f := float64(fact(n))
	pref := float64(fact(n - 1))
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j {
				x := p[i][0] - p[j][0]
				y := p[i][1] - p[j][1]
				out += math.Sqrt(float64(x*x+y*y)) * pref
			}
		}
	}
	return out / f
}

func fact(x int) int {
	fact := 1
	for i := 1; i <= x; i++ {
		fact *= i
	}
	return fact
}
func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	p := make([][2]int, n)
	for i := range p {
		fmt.Fscan(r, &p[i][0], &p[i][1])
	}
	fmt.Printf("%.10f", solution(n, p))
}
