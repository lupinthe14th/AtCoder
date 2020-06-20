package main

import (
	"fmt"
)

func solution(s string) string {
	const KEY = "keyence"
	if s == KEY {
		return "YES"
	}
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if KEY == s[:i]+s[j:] {
				return "YES"
			}
		}
	}
	return "NO"
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(solution(s))
}
