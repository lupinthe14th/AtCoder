package main

import (
	"fmt"
	"strconv"
)

func solution(a, b int) int {
	out := 0
	for i := a; i <= b; i++ {
		s := strconv.Itoa(i)
		if isPalindrome(s) {
			out++
		}
	}
	return out
}

func isPalindrome(s string) bool {
	l := len(s)

	for i := 0; i < l; i++ {
		if s[i] != s[l-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(solution(a, b))
}
