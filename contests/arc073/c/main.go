package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(n, t int, ts []int) uint64 {
	var sum uint64 = 0
	for i := 1; i < n; i++ {
		d := ts[i] - ts[i-1]
		if d >= t {
			sum += uint64(t)
		} else {
			sum += uint64(d)
		}
	}
	return sum + uint64(t)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, t int
	fmt.Fscan(r, &n, &t)
	ts := make([]int, n)
	for i := range ts {
		fmt.Fscan(r, &ts[i])
	}
	fmt.Println(solution(n, t, ts))
}
