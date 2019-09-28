package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Matrix(h, w, a, b int) [][]int {
	log.Printf("h: %d", h)
	log.Printf("w: %d", w)
	log.Printf("a: %d", a)
	log.Printf("b: %d", b)
	// Slice of Slice
	dary := make([][]int, h)
	for i := range dary {
		dary[i] = make([]int, w)
	}

	for i := 0; i <= h; i++ {
		for j := 0; j <= w; j++ {
			if i+1 == b && j+1 == a {
				dary[i][j] = 0
			}
			if i+1 == b && j+1 == w-a {
				dary[i][j] = 1
			}
			if i+1 == h-b && j+1 == a {
				dary[i][j] = 1
			}
			if i+1 == h-b && j+1 == w-a {
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
	fmt.Println(matrix)
	for _, l := range matrix {
		sl := make([]string, 0, w)
		for _, s := range l {
			sl = append(sl, strconv.Itoa(s))
		}
		fmt.Println(strings.Join(sl, ""))
	}
}
