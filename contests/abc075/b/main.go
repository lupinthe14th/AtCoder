package main

import (
	"fmt"
)

func solution(h, w int, s []string) []string {
	m := make([][]byte, h)
	for i := range m {
		m[i] = make([]byte, w)
	}
	var dx = [8]int{1, 0, -1, 1, -1, 1, 0, -1}
	var dy = [8]int{-1, -1, -1, 0, 0, 1, 1, 1}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] != '#' {
				c := 0
				for d := 0; d < 8; d++ {
					ni := i + dy[d]
					nj := j + dx[d]

					if ni < 0 || h <= ni {
						continue
					}
					if nj < 0 || w <= nj {
						continue
					}
					if s[ni][nj] == '#' {
						c++
					}
				}
				m[i][j] = byte(c + '0')
			} else {
				m[i][j] = '#'
			}
		}
	}
	for i := range m {
		s[i] = string(m[i])
	}
	return s
}

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	s := make([]string, h)
	for i := range s {
		fmt.Scan(&s[i])
	}
	for _, t := range solution(h, w, s) {
		fmt.Println(t)
	}
}
