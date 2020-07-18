package main

import (
	"fmt"
)

func solution(n int, a []int) []int {

	memo := make(map[int]int)
	for i := range a {
		memo[a[i]]++
	}
	var tot int64
	for _, v := range memo {
		tot += combi2(int64(v))
	}
	out := make([]int, n)
	for i := 0; i < n; i++ {
		var ans int64 = tot
		ans -= combi2(int64(memo[a[i]]))
		memo[a[i]]--
		ans += combi2(int64(memo[a[i]]))
		memo[a[i]]++
		out[i] = int(ans)
	}
	return out
}

func combi2(x int64) int64 {
	return x * (x - 1) / 2
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	for _, t := range solution(n, a) {
		fmt.Println(t)
	}
}
