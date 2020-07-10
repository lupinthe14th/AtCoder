package main

import (
	"fmt"
)

func solution(s string) int {
	if len(s) == 0 {
		return 0
	}
	z, o := 0, 0
	for i := range s {
		if s[i] == '0' {
			z++
		} else {
			o++
		}
	}
	return 2 * min(z, o)

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(solution(s))
}
