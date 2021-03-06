package main

import (
	"bufio"
	"fmt"
	"os"
)

// SeeAlso: https://atcoder.jp/contests/abc157/submissions/10471039
func solution(n, m, k int, ab, cd [][2]int) []int {
	friends := make([]int, n)
	uf := newUnionFind(n)
	for i := 0; i < m; i++ {
		a := ab[i][0] - 1
		b := ab[i][1] - 1
		friends[a]++
		friends[b]++
		uf.unite(a, b)
	}
	blocks := make([]int, n)
	for i := 0; i < k; i++ {
		c := cd[i][0] - 1
		d := cd[i][1] - 1
		if uf.same(c, d) {
			blocks[c]++
			blocks[d]++
		}
	}
	out := make([]int, n)
	for i := 0; i < n; i++ {
		out[i] = uf.size(i) - friends[i] - blocks[i] - 1
	}
	return out
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, m, k int
	fmt.Fscan(r, &n, &m, &k)
	ab := make([][2]int, m)
	for i := range ab {
		fmt.Fscan(r, &ab[i][0], &ab[i][1])
	}
	cd := make([][2]int, k)
	for i := range cd {
		fmt.Fscan(r, &cd[i][0], &cd[i][1])
	}
	for _, t := range solution(n, m, k, ab, cd) {
		fmt.Printf("%v ", t)
	}
}

type unionFind struct {
	d []int
}

func newUnionFind(n int) unionFind {
	d := make([]int, n)
	for i := 0; i < n; i++ {
		d[i] = -1
	}
	return unionFind{d: d}
}

func (uf unionFind) find(x int) int {
	if uf.d[x] < 0 {
		return x
	}
	uf.d[x] = uf.find(uf.d[x])
	return uf.d[x]
}

func (uf unionFind) unite(x, y int) bool {
	x, y = uf.find(x), uf.find(y)
	if x == y {
		return false
	}
	if uf.d[x] > uf.d[y] {
		x, y = y, x
	}
	uf.d[x] += uf.d[y]
	uf.d[y] = x
	return true
}

func (uf unionFind) same(x, y int) bool {
	return uf.find(x) == uf.find(y)
}

func (uf unionFind) size(x int) int {
	return -uf.d[uf.find(x)]
}
