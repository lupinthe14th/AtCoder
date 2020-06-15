package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func solution(n int, a []int) int {
	sort.Ints(a)
	amax := a[n-1]

	memo := make([]int, amax+1)
	for i := 0; i < n; i++ {
		for j := 1; a[i]*j <= amax; j++ {
			if j == 1 && memo[a[i]] == 1 {
				memo[a[i]]++
				break
			}
			memo[a[i]*j]++
		}
	}
	c := 0
	for i := range a {
		if memo[a[i]] == 1 {
			c++
		}
	}
	return c
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(r, &a[i])
	}
	fmt.Println(solution(n, a))
}
