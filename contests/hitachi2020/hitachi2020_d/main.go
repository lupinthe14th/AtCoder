package main

import (
	"fmt"
	"sort"
)

type wt struct {
	a, b, c int
}

type Wts []wt

func (w Wts) Len() int {
	return len(w)
}

func (w Wts) Less(i, j int) bool {
	if w[i].c == w[j].c {
		return w[i].a < w[j].a
	}
	return w[i].c < w[j].c
}

func (w Wts) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func solution(n, t int, m []wt) int {
	sort.Sort(Wts(m))
	var T float64 = float64(t) + 0.5
	var c, w, time int
	for i := 0; i < n; i++ {
		time++
		w = m[i].a*time + m[i].b
		time += w
		if float64(time) >= T {
			break
		}
		c++
	}
	return c
}

func main() {
	var n, t int
	fmt.Scan(&n, &t)

	m := make([]wt, n)
	for i := range m {
		var a, b int
		fmt.Scan(&a, &b)
		m[i] = wt{a: a, b: b, c: a + b}
	}

	fmt.Println(solution(n, t, m))
}
