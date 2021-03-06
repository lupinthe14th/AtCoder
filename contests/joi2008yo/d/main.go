package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func solution(m, n int, p, q [][2]int) [2]int {
	sort.SliceStable(p, func(i, j int) bool {
		if p[i][0] == p[j][0] {
			return p[i][1] < p[j][1]
		}
		return p[i][0] < p[j][0]
	})
	memo := make(map[[2]int]bool, n)
	for i := range q {
		memo[q[i]] = true
	}

	p0x := p[0][0]
	p0y := p[0][1]
	var out [2]int
	for i := range q {
		dx := q[i][0] - p0x
		dy := q[i][1] - p0y
		flag := false
		for j := 1; j < m; j++ {
			x := p[j][0] + dx
			y := p[j][1] + dy
			if !memo[[2]int{x, y}] {
				flag = false
				break
			}
			flag = true
		}
		if flag {
			out = [2]int{dx, dy}
			break
		}
	}
	return out
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var m int
	fmt.Fscan(r, &m)
	p := make([][2]int, m)
	for i := range p {
		fmt.Fscan(r, &p[i][0], &p[i][1])
	}
	var n int
	fmt.Fscan(r, &n)
	q := make([][2]int, n)
	for i := range q {
		fmt.Fscan(r, &q[i][0], &q[i][1])
	}
	ans := solution(m, n, p, q)
	fmt.Printf("%v %v\n", ans[0], ans[1])
}
