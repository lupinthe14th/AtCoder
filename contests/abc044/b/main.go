package main

import (
	"fmt"
)

func solution(w string) string {
	var memo [26]int

	for _, r := range w {
		memo[r-'a']++
	}

	for i := range memo {
		if memo[i]&1 == 1 {
			return "No"
		}
	}
	return "Yes"
}

func main() {
	var w string
	fmt.Scan(&w)
	fmt.Println(solution(w))
}
