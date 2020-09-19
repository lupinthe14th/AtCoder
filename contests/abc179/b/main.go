package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(n int, d [][2]int) string {
	out, c := 0, 0
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	for i := range d {
		if d[i][0] == d[i][1] {
			c++
		} else {
			c = 0
		}
		out = max(out, c)
	}
	if out >= 3 {
		return "Yes"
	}
	return "No"
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	d := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &d[i][0], &d[i][1])
	}
	fmt.Println(solution(n, d))
}
