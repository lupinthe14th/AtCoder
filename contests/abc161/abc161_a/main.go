package main

import (
	"fmt"
)

func solution(x, y, z int) (a, b, c int) {
	a = y
	b = x
	a = z
	c = y
	return a, b, c
}

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	a, b, c := solution(x, y, z)
	fmt.Printf("%d %d %d\n", a, b, c)
}
