package main

import (
	"fmt"
)

func roundOne(a, b int) int {
	var c int
	switch ans := a * b; ans {
	case 3:
		c = 2
	case 2:
		c = 3
	case 6:
		c = 1
	}
	return c
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(roundOne(a, b))
}
