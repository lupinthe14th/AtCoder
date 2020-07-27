package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(n, m int, g [][3]int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, i+2)
	}
	for i := 0; i < m; i++ {
		x := g[i][0]
		y := g[i][1]
		z := g[i][2] + 1
		dp[x][y] = z
	}
	out := 0
	for i := 1; i < n+1; i++ {
		for j := 1; j < i+1; j++ {
			dp[i][j] = max(dp[i][j], max(dp[i-1][j], dp[i-1][j-1])-1)
			if dp[i][j] != 0 {
				out++
			}
		}
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
	var n, m int
	fmt.Fscan(r, &n, &m)
	g := make([][3]int, m)
	for i := range g {
		fmt.Fscan(r, &g[i][0], &g[i][1], &g[i][2])
	}
	fmt.Println(solution(n, m, g))
}
