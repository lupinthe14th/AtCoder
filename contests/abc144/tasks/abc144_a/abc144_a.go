package main

import (
	"fmt"
)

func multiplicationTable(x, y int) int {
	if x > 9 || y > 9 {
		return -1
	}
	return x * y
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(multiplicationTable(a, b))
}
