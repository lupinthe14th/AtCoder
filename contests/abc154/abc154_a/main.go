package main

import (
	"fmt"
)

func solution(s, t, u string, a, b int) (m, n int) {
	switch {
	case s == u:
		a--
	case t == u:
		b--
	}
	return a, b
}

func main() {
	var s, t, u string
	var a, b int
	fmt.Scan(&s, &t, &a, &b, &u)
	fmt.Println(solution(s, t, u, a, b))
}
