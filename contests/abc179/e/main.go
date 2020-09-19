package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(n, x, m int) int {

	a := make([]int, 0)
	memo := make(map[int]int, m+1)
	f := func(a, b int) int {
		return a % b
	}
	a = append(a, x)
	memo[x] = 0
	cur := x

	start, last := 0, 0
	for i := 0; i < m; i++ {
		cur = f(cur*cur, m)
		a = append(a, cur)
		idx, ok := memo[cur]
		if ok {
			start = idx
			last = i + 1
			break
		} else {
			memo[cur] = i + 1
		}
	}

	// 繰り返しが始まるまでの和
	bsum := 0
	for i := 0; i < start; i++ {
		bsum += a[i]
	}
	// 繰り返し部分の和
	asum := 0
	for i := start; i < last; i++ {
		asum += a[i]
	}
	// 繰り返す部分の回数
	c := (n - start) / (last - start)
	// 端数
	mod := (n - start) % (last - start)
	csum := 0
	for i := start; i < start+mod; i++ {
		csum += a[i]
	}
	return bsum + asum*c + csum
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, x, m int
	fmt.Fscan(r, &n, &x, &m)
	fmt.Println(solution(n, x, m))
}
