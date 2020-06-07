package main

import (
	"fmt"
)

func solution(h, w int, a []string) []string {
	out := make([]string, h+2)
	for i := range out {
		buf := make([]byte, 0, w+2)
		for j := 0; j < w+2; j++ {
			if i > 0 && i < h+1 && j > 0 && j < w+1 {
				buf = append(buf, a[i-1][j-1])
			} else {
				buf = append(buf, '#')
			}
		}
		out[i] = string(buf)
	}
	return out
}

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	a := make([]string, h)
	for i := range a {
		fmt.Scan(&a[i])
	}
	for _, t := range solution(h, w, a) {
		fmt.Println(t)
	}
}
