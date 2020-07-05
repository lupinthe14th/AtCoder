package main

import (
	"fmt"
)

func solution(h, w, k int, c [][]byte) int {
	out := 0
	for R := 0; R < 1<<uint(h)-1; R++ {
		for C := 0; C < 1<<uint(w)-1; C++ {
			b := 0
			for i := range c {
				for j := range c[i] {
					if R>>uint(i)&1 == 0 && C>>uint(j)&1 == 0 && c[i][j] == '#' {
						b++
					}
				}
			}
			if b == k {
				out++
			}
		}
	}
	return out
}

func re(t, c [][]byte) [][]byte {
	t = make([][]byte, len(c)+1)
	for i := range t {
		t[i] = make([]byte, len(c[0])+1)
	}
	for i := range c {
		for j := range c[i] {
			t[i+1][j+1] = c[i][j]
		}
	}
	return t
}
func count(c [][]byte) int {
	cnt := 0
	for i := range c {
		for j := range c[i] {
			if c[i][j] == '#' {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	var h, w, k int
	fmt.Scan(&h, &w, &k)
	c := make([][]byte, h)
	for i := range c {
		var t string
		c[i] = make([]byte, w)
		fmt.Scan(&t)
		for j := 0; j < w; j++ {
			c[i][j] = t[j]
		}
	}
	fmt.Println(solution(h, w, k, c))
}
