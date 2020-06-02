package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	var b int
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		if i > 0 {
			switch {
			case a == b:
				fmt.Println("stay")
			case a > b:
				fmt.Printf("up %d\n", a-b)
			case a < b:
				fmt.Printf("down %d\n", b-a)
			}
		}
		b = a
	}
}
