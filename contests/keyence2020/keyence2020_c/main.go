package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solution(n, k, s int) []int {
	ans := make([]int, n)
	for i := 0; i < k; i++ {
		ans[i] = s
	}
	if s == 1e9 {
		for i := k; i < n; i++ {
			ans[i] = 1
		}
	} else {
		for i := k; i < n; i++ {
			ans[i] = s + 1
		}
	}

	return ans
}

func main() {
	var n, k, s int
	fmt.Scan(&n, &k, &s)
	ans := solution(n, k, s)

	l := len(ans)
	strs := make([]string, l)

	for i := range ans {
		strs[i] = strconv.Itoa(ans[i])
	}

	fmt.Println(strings.Join(strs, " "))
}
