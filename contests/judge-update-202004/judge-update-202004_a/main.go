package main

import (
	"fmt"
)

func solution(s, l, r int) int {
	if l <= s && s <= r {
		return s
	}
	for s < l {
		s++
	}
	for s > r {
		s--
	}
	return s
}

func main() {
	var s, l, r int
	fmt.Scan(&s, &l, &r)
	fmt.Println(solution(s, l, r))
}
