package main

import (
	"fmt"
)

func solution(s string) string {
	var a [26]int
	for _, r := range s {
		if a[r-'a'] != 0 {
			return "no"
		}
		a[r-'a']++
	}
	return "yes"
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(solution(s))
}
