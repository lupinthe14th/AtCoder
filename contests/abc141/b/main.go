package main

import (
	"fmt"
)

func solution(s string) string {
	for i, r := range s {
		if (i-1)%2 == 0 {
			if r == 'R' {
				return "No"
			}
		} else {
			if r == 'L' {
				return "No"
			}
		}
	}
	return "Yes"
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(solution(s))
}
