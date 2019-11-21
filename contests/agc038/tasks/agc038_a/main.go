package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Matrix(h, w, a, b int) [][]int {
	// Slice of Slice
	dary := make([][]int, h)
	for i := range dary {
		dary[i] = make([]int, w)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < a; j++ {
			if i < b {
				dary[i][j] = 0
			} else {
				dary[i][j] = 1
			}
		}
		for j := a; j < w; j++ {
			if i < b {
				dary[i][j] = 1
			} else {
				dary[i][j] = 0
			}
		}
	}
	return dary
}

func main() {
	var h, w, a, b int
	fmt.Scan(&h, &w, &a, &b)
	matrix := Matrix(h, w, a, b)
	for _, l := range matrix {
		sl := make([]string, w)
		for _, s := range l {
			sl = append(sl, strconv.Itoa(s))
		}
		fmt.Println(strings.Join(sl, ""))
	}
}
