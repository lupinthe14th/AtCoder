package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(n int) int {
	if n == 0 {
		return 1
	}
	return 0
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	fmt.Println(solution(n))
}
