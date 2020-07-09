package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(n, m int, a, b []int) (string, []int) {
	const MAX = 1 << 31
	g := make([][]int, n)
	for i := 0; i < m; i++ {
		v, w := a[i]-1, b[i]-1
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	queue := make([]int, 0, n)

	dist := make([]int, n)
	for i := range dist {
		dist[i] = MAX
	}

	prev := make([]int, n)
	for i := range prev {
		prev[i] = -1
	}

	queue = append(queue, 0)
	dist[0] = 0
	prev[0] = 0

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		for _, u := range g[v] {
			if dist[u] != MAX {
				continue
			}
			dist[u] = dist[v] + 1
			prev[u] = v
			queue = append(queue, u)
		}

	}

	for _, t := range prev {
		if t == -1 {
			return "No", []int{}
		}
	}
	for i := 1; i < n; i++ {
		prev[i]++
	}
	return "Yes", prev[1:]
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(r, &n, &m)
	a := make([]int, 0, m)
	b := make([]int, 0, m)
	for i := 0; i < m; i++ {
		var ta, tb int
		fmt.Fscan(r, &ta, &tb)
		a = append(a, ta)
		b = append(b, tb)

	}
	ans, s := solution(n, m, a, b)
	fmt.Println(ans)
	if s != nil {
		for _, t := range s {
			fmt.Println(t)
		}
	}
}
