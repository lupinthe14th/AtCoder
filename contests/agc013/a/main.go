package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(n int, a []int) int {
	c := 0
	rise := 0 // -1: down, 1: up
	for i := 1; i < n; i++ {
		d := a[i] - a[i-1]
		switch {
		case d > 0 && rise == -1:
			c++
			rise = 0
		case d < 0 && rise == 1:
			c++
			rise = 0
		case d > 0:
			rise = 1
		case d < 0:
			rise = -1
		}
	}
	return c + 1
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
