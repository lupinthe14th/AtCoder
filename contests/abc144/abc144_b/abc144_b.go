package main

import (
	"fmt"
)

func isMultiplicationTable(x int) string {
	multiplicationTableNumbers := make(map[int]bool)

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if !multiplicationTableNumbers[i*j] {
				multiplicationTableNumbers[i*j] = true
			}
		}
	}
	if multiplicationTableNumbers[x] {
		return "Yes"
	}
	return "No"
}

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println(isMultiplicationTable(a))
}
