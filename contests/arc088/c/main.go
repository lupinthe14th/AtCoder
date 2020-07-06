package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(x, y int) int {
	c := 0
	for cur := x; cur <= y; cur <<= 1 {
		c++
	}
	return c
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var x, y int
	fmt.Fscan(r, &x, &y)
	fmt.Println(solution(x, y))
}
