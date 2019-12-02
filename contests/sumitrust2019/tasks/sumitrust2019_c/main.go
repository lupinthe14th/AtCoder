package main

import (
	"fmt"
)

// 品物の個数を全探索
func oneOOToOneOFive(x int) bool {
	for c := 1; c <= x; c++ {
		if 100*c <= x && x <= 105*c {
			return true
		}
	}
	return false
}

func main() {
	var x int
	fmt.Scan(&x)
	if oneOOToOneOFive(x) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
