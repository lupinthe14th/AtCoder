package main

import (
	"fmt"
)

func roundOne(a, b int) int {
	return a ^ b
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(roundOne(a, b))
}
