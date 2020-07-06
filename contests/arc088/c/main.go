package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func solution(x, y int) int {
	memo := make([]int, 0)
	hi := 0
	for i := 1; hi <= y; i++ {
		hi = x * int(math.Pow(2, float64(i)-1))
		memo = append(memo, hi)
	}
	return len(memo) - 1
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var x, y int
	fmt.Fscan(r, &x, &y)
	fmt.Println(solution(x, y))
}
