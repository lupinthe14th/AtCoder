package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func solution(n int, a []int) int {
	sort.Ints(a)
	memo := make([][3]int, n)
	for i := 0; i < n; i++ {
		memo[i][0] = a[i]
	}

	j := 0
	for i := 3*n - 1; i > n; i -= 2 {
		memo[j][2] = a[i]
		memo[j][1] = a[i-1]
		j++
	}

	out := 0
	for i := range memo {
		out += memo[i][1]
	}
	return out
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	a := make([]int, 3*n)
	for i := range a {
		fmt.Fscan(r, &a[i])
	}
	fmt.Println(solution(n, a))
}
