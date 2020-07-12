package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(n int, a []int) int {
	curI := make([]int, n)
	curD := make([]int, n)
	curI[0], curD[n-1] = a[0], a[n-1]
	for i := 1; i < n; i++ {
		curI[i] = a[i] + curI[i-1]
	}
	for i := n - 2; i >= 0; i-- {
		curD[i] = a[i] + curD[i+1]
	}
	out := 1 << 31
	for i := 0; i < n-1; i++ {
		out = min(out, abs(curI[i]-curD[i+1]))
	}
	return out
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(r, &a[i])
	}
	fmt.Println(solution(n, a))
}
