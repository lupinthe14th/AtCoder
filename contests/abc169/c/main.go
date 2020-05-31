package main

import (
	"fmt"
	"strconv"
)

func solution(a int, b string) int {
	x := make([]byte, 0, len(b))
	for i := range b {
		if b[i] != '.' {
			x = append(x, b[i])
		}
	}
	bn, _ := strconv.Atoi(string(x))

	return a * bn / 100
}

func main() {
	var a int
	var b string
	fmt.Scan(&a, &b)
	fmt.Println(solution(a, b))
}
