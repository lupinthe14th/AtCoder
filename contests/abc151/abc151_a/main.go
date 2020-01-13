package main

import (
	"fmt"
)

func solution(c string) string {
	var ans string
	for _, r := range c {
		ans = string(r + 1)
	}
	return ans
}

func main() {
	var c string
	fmt.Scan(&c)
	fmt.Println(solution(c))
}
