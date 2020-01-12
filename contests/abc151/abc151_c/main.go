package main

import (
	"fmt"
)

func solution(n, m int, r []result) ans {
	memo := make(map[int]int, n)
	var c, p int
	for i := 0; i < m; i++ {
		if memo[r[i].p] == -1 {
			continue
		} else if r[i].s == "AC" {
			c++
			p += memo[r[i].p]
			memo[r[i].p] = -1
		} else {
			memo[r[i].p]++
		}
	}
	return ans{c: c, p: p}
}

type result struct {
	p int
	s string
}

type ans struct {
	c, p int
}

func main() {
	var n, m, p int
	var s string
	r := make([]result, 0, m)
	fmt.Scan(&n, &m)
	for i := 0; i < m; i++ {
		fmt.Scan(&p, &s)
		r = append(r, result{p: p, s: s})
	}
	ans := solution(n, m, r)
	fmt.Printf("%d %d", ans.c, ans.p)
}
