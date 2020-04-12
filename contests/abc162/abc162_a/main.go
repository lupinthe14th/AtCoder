package main

import (
	"fmt"
)

func solution(n int) string {
	for n > 0 {
		if n%10 == 7 {
			return "Yes"
		}
		n /= 10
	}
	return "No"
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Printf(solution(n))
}
