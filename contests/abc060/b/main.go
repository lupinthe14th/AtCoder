package main

import (
	"fmt"
)

func solution(a, b, c int) string {
	for i := 1; i < b+1; i++ {
		if i*a%b == c {
			return "YES"
		}
	}
	return "NO"
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(solution(a, b, c))
}
