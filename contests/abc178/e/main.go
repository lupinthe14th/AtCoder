package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(n int, p [][2]int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	zmax, zmin, wmax, wmin := -1<<31, 1<<31, -1<<31, 1<<31
	for i := range p {
		z := p[i][0] + p[i][1]
		w := p[i][0] - p[i][1]
		zmax = max(zmax, z)
		zmin = min(zmin, z)
		wmax = max(wmax, w)
		wmin = min(wmin, w)
	}
	return max(zmax-zmin, wmax-wmin)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	p := make([][2]int, n)
	for i := range p {
		fmt.Fscan(r, &p[i][0], &p[i][1])
	}
	fmt.Println(solution(n, p))
}
