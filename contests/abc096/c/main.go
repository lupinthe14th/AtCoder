package main

import (
	"fmt"
	"strings"
)

func solution(h, w int, s [][]string) string {
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == "#" {
				if i > 0 && s[i-1][j] == "#" {
					continue
				}
				if i < h-1 && s[i+1][j] == "#" {
					continue
				}
				if j > 0 && s[i][j-1] == "#" {
					continue
				}
				if j < w-1 && s[i][j+1] == "#" {
					continue
				}
				return "No"
			}
		}
	}
	return "Yes"
}

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	s := make([][]string, h)
	for i := range s {
		var t string
		fmt.Scan(&t)
		s[i] = strings.Split(t, "")
	}
	fmt.Println(solution(h, w, s))
}
