package main

import (
	"fmt"
)

func solution(n, a, b int, s string) []string {
	ans := make([]string, 0, n)

	in, ext := 0, 0
	for _, r := range s {
		res := "No"
		switch r {
		case 'a':
			if a+b > in+ext {
				res = "Yes"
				in++
			}
			ans = append(ans, res)
		case 'b':
			if a+b > in+ext && b >= ext+1 {
				res = "Yes"
				ext++
			}
			ans = append(ans, res)
		default:
			ans = append(ans, "No")
		}
	}
	return ans
}

func main() {
	var n, a, b int
	var s string
	fmt.Scan(&n, &a, &b, &s)
	for _, text := range solution(n, a, b, s) {
		fmt.Println(text)
	}
}
