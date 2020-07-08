package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(n, m, q int, s [][]int) []int {
	type user struct {
		id     int
		scores []int
		score  int
	}
	u := make([]user, n)
	for i := range u {
		u[i].scores = make([]int, m)
	}
	memo := make([]int, m)
	out := make([]int, 0, n)
	for i := range s {
		cid := s[i][0] // case id: 1 || 2
		switch cid {
		case 1:
			uid := s[i][1] - 1
			out = append(out, u[uid].score)
		case 2:
			uid := s[i][1] - 1 // uid
			pid := s[i][2] - 1 // Problem id
			memo[pid]++
			score := n - memo[pid]
			u[uid].scores[pid] = score
			u[uid].score += score
			for k := range u {
				if k != uid && u[k].scores[pid] > 0 {
					t := u[k].scores[pid]
					u[k].scores[pid] = score
					u[k].score -= t - score
				}
			}
		}
	}
	return out
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, m, q int
	fmt.Fscan(r, &n, &m, &q)
	s := make([][]int, q)
	for i := range s {
		var t int
		fmt.Fscan(r, &t)
		switch t {
		case 1:
			s[i] = make([]int, 2)
			s[i][0] = t
			fmt.Fscan(r, &s[i][1])
		case 2:
			s[i] = make([]int, 3)
			s[i][0] = t
			fmt.Fscan(r, &s[i][1], &s[i][2])
		}
	}
	for _, t := range solution(n, m, q, s) {
		fmt.Println(t)
	}
}
