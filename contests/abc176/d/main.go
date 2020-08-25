package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// using 01-BFS
// seealso: https://atcoder.jp/contests/abc176/submissions/16224839
func solution(h, w int, c, d [2]int, s []string) int {

	c[0], c[1] = c[0]-1, c[1]-1
	d[0], d[1] = d[0]-1, d[1]-1

	dx := [4]int{-1, 0, 1, 0}
	dy := [4]int{0, -1, 0, 1}

	const INF = 1e9
	dist := make([][]int, h)
	for i := 0; i < h; i++ {
		dist[i] = make([]int, w)
		for j := 0; j < w; j++ {
			dist[i][j] = INF
		}
	}
	// dequeを使う
	// cost: 0を使う場合は前からappend
	// cost: 1を使う場合は後からappend
	deque := &queues{}
	heap.Push(deque, queue{x: c[0], y: c[1], z: 0})
	for deque.Len() > 0 {
		cur := heap.Pop(deque).(queue)
		x := cur.x
		y := cur.y
		z := cur.z
		for i := 0; i < 4; i++ {
			nx := x + dx[i]
			ny := y + dy[i]
			if nx < 0 || nx >= h {
				continue
			}
			if ny < 0 || ny >= w {
				continue
			}
			if s[nx][ny] == '#' {
				continue
			}
			if dist[nx][ny] <= z {
				continue
			}
			dist[nx][ny] = z
			heap.Push(deque, queue{x: nx, y: ny, z: z})
		}
		// warp
		for i := -2; i <= 2; i++ {
			for j := -2; j <= 2; j++ {
				if i == 0 && j == 0 {
					continue
				}
				nx := x + i
				ny := y + j
				if nx < 0 || nx >= h {
					continue
				}
				if ny < 0 || ny >= w {
					continue
				}
				if s[nx][ny] == '#' {
					continue
				}
				if dist[nx][ny] <= z+1 {
					continue
				}
				dist[nx][ny] = z + 1
				heap.Push(deque, queue{x: nx, y: ny, z: z + 1})
			}
		}
	}
	out := dist[d[0]][d[1]]
	if out == INF {
		return -1
	}
	return out
}

// An deque is a min-heap of ints.
type queue struct {
	x, y, z int
}

type queues []queue

func (h queues) Len() int           { return len(h) }
func (h queues) Less(i, j int) bool { return h[i].z < h[j].z }
func (h queues) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *queues) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(queue))
}

func (h *queues) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var h, w int
	fmt.Fscan(r, &h, &w)
	var c [2]int
	fmt.Fscan(r, &c[0], &c[1])
	var d [2]int
	fmt.Fscan(r, &d[0], &d[1])
	s := make([]string, h)
	for i := 0; i < h; i++ {
		fmt.Fscan(r, &s[i])
	}
	fmt.Println(solution(h, w, c, d, s))
}
