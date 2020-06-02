package main

import (
	"fmt"
)

func solution(s string) int {
	l := len(s)
	i := 0
	for i = l - 1; i >= 0; i-- {
		if isEvenString(s[:i]) {
			break
		}
	}
	return i
}

func isEvenString(s string) bool {
	l := len(s)
	if l&1 == 1 {
		return false
	}
	if s[0:l/2] == s[l/2:l] {
		return true
	}
	return false
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(solution(s))
}
