package main

import (
	"fmt"
	"strconv"
)

func solution(n int) string {
	t := strconv.Itoa(n)
	ans := ""

	switch t[len(t)-1] {
	case '2', '4', '5', '7', '9':
		ans = "hon"
	case '0', '1', '6', '8':
		ans = "pon"
	case '3':
		ans = "bon"
	}
	return ans
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(solution(n))
}
