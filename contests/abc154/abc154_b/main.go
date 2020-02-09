package main

import (
	"fmt"
)

func solution(s string) string {
	l := len(s)
	b := make([]byte, l)
	for i := 0; i < l; i++ {
		b[i] = byte('x')
	}
	return string(b)
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(solution(s))
}
