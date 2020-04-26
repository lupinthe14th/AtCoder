package main

import (
	"fmt"
)

func solution(a, b, c, d int) string {
	//t: ab
	//a: cd
	var ans string
	i := 0
	for {
		if (i & 1) == 0 {
			c -= b
		} else {
			a -= d
		}
		if a <= 0 {
			ans = "No"
			break
		}
		if c <= 0 {
			ans = "Yes"
			break
		}
		i++
	}
	return ans
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	fmt.Println(solution(a, b, c, d))
}
