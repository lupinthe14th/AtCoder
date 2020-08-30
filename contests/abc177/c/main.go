package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(n int, a []int) int {
	const MOD = 1e9 + 7
	out := 0
	cum := make([]int, n+1)
	for i := 0; i < n; i++ {
		cum[i+1] = a[i] + cum[i]
	}
	for i := 0; i < n; i++ {
		sum := (cum[n] - cum[i+1]) % MOD
		out += a[i] * sum
		out %= MOD
	}
	return out
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
