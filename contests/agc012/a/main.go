package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func solution(n int, a []int) int {
	sort.Ints(a)
	out := 0
	for i := n; i < 3*n; i += 2 {
		out += a[i]
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
