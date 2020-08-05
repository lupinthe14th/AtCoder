package main

import (
	"fmt"
)

func solution(n, m int, xy [][2]int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	g := make([][]bool, n)
	for i := range g {
		g[i] = make([]bool, n)
	}
	for i := range xy {
		x := xy[i][0] - 1
		y := xy[i][1] - 1
		g[x][y] = true
		g[y][x] = true
	}
	out := 1
	for s := 0; s < 1<<uint(n); s++ {
		c := 0
		ok := true
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				if s>>uint(i)&1 == 1 && s>>uint(j)&1 == 1 && !g[i][j] {
					ok = false
					break
				}
			}
		}
		if ok {
			for i := 0; i < n; i++ {
				if s>>uint(i)&1 == 1 {
					c++
				}
			}
			out = max(out, c)
		}
	}
	return out
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	xy := make([][2]int, m)
	for i := range xy {
		fmt.Scan(&xy[i][0], &xy[i][1])
	}
	fmt.Println(solution(n, m, xy))
}
