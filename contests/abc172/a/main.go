package main

import (
	"fmt"
)

func solution(a int) int {
	return a + a*a + a*a*a
}

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println(solution(a))
}
