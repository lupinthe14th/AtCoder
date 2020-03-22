package main

import (
	"fmt"
)

func solution(s string) string {
	if !isPalindrome(s) {
		return "No"
	}
	if !isPalindrome(s[:(len(s)-1)/2]) {
		return "No"
	}
	if !isPalindrome(s[-1+(len(s)+3)/2:]) {
		return "No"
	}
	return "Yes"
}

func isPalindrome(s string) bool {
	r := make([]rune, len(s))

	for i, v := range s {
		r[len(s)-1-i] = v
	}

	if s == string(r) {
		return true
	}

	return false
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(solution(s))
}
