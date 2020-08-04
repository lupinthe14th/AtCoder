package main

import (
	"fmt"
)

func solution(n, m int, k, p []int, s [][]int) int {

	out := 0

	for a := 0; a < 1<<uint(n); a++ {
		var ok bool = true
		for i := 0; i < m; i++ {
			c := 0
			for _, j := range s[i] {
				if a>>uint(j-1)&1 == 0 {
					c++
				}
			}
			c %= 2
			if c != p[i] {
				ok = false
			}
		}
		if ok {
			out++
		}
	}
	return out
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	k := make([]int, m)
	s := make([][]int, m)
	for i := range k {
		fmt.Scan(&k[i])
		s[i] = make([]int, k[i])
		for j := range s[i] {
			fmt.Scan(&s[i][j])
		}
	}
	p := make([]int, m)
	for i := range p {
		fmt.Scan(&p[i])
	}
	fmt.Println(solution(n, m, k, p, s))
}
