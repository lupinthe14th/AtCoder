package main

import (
	"fmt"
	"log"
)

func solution(n, m, x int, a [][]int, c []int) int {
	s := make([][]int, n+1)
	for i := range s {
		s[i] = make([]int, m)
	}
	for i := 1; i <= n; i++ {
		for j := 0; j < m; j++ {
			s[i][j] = s[i-1][j] + a[i-1][j]
		}
	}
	log.Print(s)

	mi := 1 << 31
	for j := 0; j < m; j++ {
		mi = min(mi, s[n][j])
	}
	if mi < x {
		return -1
	}

	mi = 1 << 31
	b, e := 0, 0
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < m; j++ {
			mi = min(mi, s[i][j])
		}
		if mi >= x {
			e = i
		}
	}
	log.Print(e)
	for b = 0; b < n; b++ {
		log.Print(s[b])
		for j := 0; j < m; j++ {
			mi = min(mi, s[e][j]-s[b][j])
		}
		if mi >= x {
			break
		}
	}
	log.Print(b)

	ans := 0
	for k := b; k <= e; k++ {
		ans += c[k]
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	var n, m, x int
	fmt.Scan(&n, &m, &x)

	c := make([]int, 0, n)
	a := make([][]int, n)

	for i := 0; i < n; i++ {
		tmp := 0
		fmt.Scan(&tmp)
		c = append(c, tmp)

		a[i] = make([]int, 0, m)
		for j := 0; j < m; j++ {
			tmp := 0
			fmt.Scan(&tmp)
			a[i] = append(a[i], tmp)
		}
	}
	fmt.Println(solution(n, m, x, a, c))
}
