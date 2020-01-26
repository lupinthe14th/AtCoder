package main

import (
	"fmt"
)

func solution(h, a int) int {
	var i int = 1
	for h > a*i {
		i++
	}
	return i
}

func main() {
	var h, a int
	fmt.Scan(&h, &a)
	fmt.Println(solution(h, a))
}
