package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(n, m, k int, a, b []int) int {
	cumA := make([]int, n+1)
	cumB := make([]int, m+1)

	for i := range a {
		cumA[i+1] = cumA[i] + a[i]
	}
	for i := range b {
		cumB[i+1] = cumB[i] + b[i]
	}

	out, j := 0, m

	for i := range cumA {
		if cumA[i] > k {
			break
		}
		for cumB[j] > k-cumA[i] {
			j--
		}
		out = max(out, i+j)
	}
	return out
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, m, k int
	fmt.Fscan(r, &n, &m, &k)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(r, &a[i])
	}
	b := make([]int, m)
	for i := range b {
		fmt.Fscan(r, &b[i])
	}
	fmt.Println(solution(n, m, k, a, b))
}
