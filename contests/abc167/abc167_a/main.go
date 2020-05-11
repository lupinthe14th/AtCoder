package main

import (
	"fmt"
)

func solution(s, t string) string {
	if s == t[:len(t)-1] {
		return "Yes"
	}
	return "No"
}

func main() {
	var s, t string
	fmt.Scan(&s, &t)
	fmt.Println(solution(s, t))
}
