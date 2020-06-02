package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	seen := make(map[int]bool, n)
	var x, y int
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		if !seen[a] {
			seen[a] = true
		} else {
			y = a
		}
	}
	for i := 0; i < n; i++ {
		if !seen[i+1] {
			x = i + 1
			break
		}
	}
	if x == 0 && y == 0 {
		fmt.Println("Correct")
	} else {
		fmt.Printf("%d %d\n", y, x)
	}
}
