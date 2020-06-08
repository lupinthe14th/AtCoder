package main

import (
	"fmt"
)

func solution(a, b string) string {
	l := max(len(a), len(b))
	a = helper(a, l)
	b = helper(b, l)
	if a > b {
		return "GREATER"
	}

	if a == b {
		return "EQUAL"
	}
	return "LESS"
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func helper(s string, l int) string {
	sl := len(s)
	if sl < l {
		buf := make([]byte, 0, l)
		for i := 0; i < l; i++ {
			if i < l-sl {
				buf = append(buf, '0')
			} else {
				buf = append(buf, s[i-l+sl])
			}
		}
		return string(buf)
	}
	return s
}

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	fmt.Println(solution(a, b))
}
