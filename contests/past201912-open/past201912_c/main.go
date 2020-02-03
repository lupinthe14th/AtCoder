package main

import (
	"fmt"
	"sort"
)

func solution(a, b, c, d, e, f int) int {
	ans := make([]int, 6)
	ans[0] = a
	ans[1] = b
	ans[2] = c
	ans[3] = d
	ans[4] = e
	ans[5] = f
	sort.Ints(ans)
	return ans[3]
}

func main() {
	var a, b, c, d, e, f int
	fmt.Scan(&a, &b, &c, &d, &e, &f)
	fmt.Println(solution(a, b, c, d, e, f))
}
