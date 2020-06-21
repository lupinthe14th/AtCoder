package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(n, t int, ts []int) int {
	var sum int = 0
	for i := 1; i < n; i++ {
		d := ts[i] - ts[i-1]
		if d >= t {
			sum += t
		} else {
			sum += d
		}
	}
	return sum + t
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
