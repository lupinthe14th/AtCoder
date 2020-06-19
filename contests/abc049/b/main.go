package main

import (
	"fmt"
)

func solution(h, w int, c []string) []string {
	out := make([]string, 2*h)

	for i := 0; i < h; i++ {
		out[2*i] = c[i]
		out[2*i+1] = c[i]
	}
	return out
}

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	c := make([]string, h)
	for i := range c {
		fmt.Scan(&c[i])
	}
	for _, t := range solution(h, w, c) {
		fmt.Println(t)
	}
}
