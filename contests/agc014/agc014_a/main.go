package main

import (
	"fmt"
)

func solution(a, b, c int) int {

	count := 0
	for a&1 == 0 && b&1 == 0 && c&1 == 0 {
		at := a / 2
		bt := b / 2
		ct := c / 2
		a = bt + ct
		b = at + ct
		c = at + bt
		count++
		if a == b && b == c && c == a {
			return -1
		}
	}
	return count
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(solution(a, b, c))
}
