package main

import (
	"fmt"
)

func solution(h, w, k int, c [][]byte) int {
	out := 0
	for row := 0; row < 1<<uint(h); row++ {
		for col := 0; col < 1<<uint(w); col++ {
			b := 0
			for i := 0; i < h; i++ {
				for j := 0; j < w; j++ {
					if row>>uint(i)&1 == 0 && col>>uint(j)&1 == 0 && c[i][j] == '#' {
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
