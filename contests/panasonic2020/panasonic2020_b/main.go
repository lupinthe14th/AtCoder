package main

import (
	"fmt"
)

func solution(h, w int) int {
	lo, hi := lohi(h, w)
	var s int
	if lo%2 == 0 {
		s = (lo * lo) / 2
	} else {
		s = (lo*lo)/2 + 1
	}

	d := hi / lo

	r := ((hi - d*lo) * lo) / 2
	return s*d + r
}

func lohi(x, y int) (lo, hi int) {
	if x < y {
		return x, y
	}
	return y, x
}

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	fmt.Println(solution(h, w))
}
