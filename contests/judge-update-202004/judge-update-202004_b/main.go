package main

import (
	"fmt"
	"sort"
)

func solution(n int, r, b []int) []int {
	ans := make([]int, 0, n)
	if len(r) != 0 {
		sort.Ints(r)
		ans = append(ans, r...)
	}
	if len(b) != 0 {
		sort.Ints(b)
		ans = append(ans, b...)
	}
	return ans
}

func main() {
	var n int
	fmt.Scan(&n)

	r := make([]int, 0, n)
	b := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var x int
		var c string
		fmt.Scan(&x, &c)
		switch {
		case c == "R":
			r = append(r, x)
		case c == "B":
			b = append(b, x)
		}
	}
	for _, v := range solution(n, r, b) {
		fmt.Println(v)
	}
}
