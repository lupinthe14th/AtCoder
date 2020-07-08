package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(n, m, q int, s [][]int) []int {
	type user struct {
		collects []int
	}
	users := make([]user, n)
	for i := range users {
		users[i].collects = make([]int, 0, m)
	}
	problems := make([]int, m)
	for i := range problems {
		problems[i] = n
	}
	out := make([]int, 0, n)
	for i := range s {
		cid := s[i][0] // case id: 1 || 2
		switch cid {
		case 1:
			uid := s[i][1] - 1
			score := 0
			for _, v := range users[uid].collects {
				score += problems[v]
			}
			out = append(out, score)
		case 2:
			uid := s[i][1] - 1 // uid
			pid := s[i][2] - 1 // Problem id
			users[uid].collects = append(users[uid].collects, pid)
			problems[pid]--
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
