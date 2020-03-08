package main

import (
	"fmt"
	"log"
)

func solution(a, b int) int {
	ans := 1 << 31
	for i := 1250; i > 0; i-- {
		if int(float64(i)*0.08) == a && int(float64(i)*0.1) == b {
			log.Print(i)
			ans = min(ans, i)
		}
	}
	if ans != 1<<31 {
		return ans
	}
	return -1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(solution(a, b))
}
