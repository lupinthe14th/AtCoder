package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(n, m int, t, p, x []int) []int {
	out := make([]int, m)
	for i := 0; i < m; i++ {
		for j, num := range t {
			if p[i] == j+1 {
				out[i] += x[i]
			} else {
				out[i] += num
			}
		}
	}
	return out
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(r, &n)
	t := make([]int, n)
	for i := range t {
		fmt.Fscan(r, &t[i])
	}
	fmt.Fscan(r, &m)
	p := make([]int, m)
	x := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &p[i], &x[i])
	}
	for _, num := range solution(n, m, t, p, x) {
		fmt.Println(num)
	}
}
