package main

import (
	"fmt"
)

func solution(n, k, q int, a []int) []string {
	p := make([]int, n)
	for i := 0; i < q; i++ {
		p[a[i]-1]++
	}

	out := make([]string, 0, n)
	for i := range p {
		if k-q+p[i] > 0 {
			out = append(out, "Yes")
		} else {
			out = append(out, "No")
		}
	}

	return out
}

func main() {
	var n, k, q int
	fmt.Scan(&n, &k, &q)
	a := make([]int, 0, q)
	for i := 0; i < q; i++ {
		t := 0
		fmt.Scan(&t)
		a = append(a, t)
	}
	for _, s := range solution(n, k, q, a) {
		fmt.Println(s)
	}
}
