package main

import (
	"fmt"
)

// 品物の個数を決めうち
func solution(x int) bool {
	for c := 1; c <= x/100; c++ {
		if 100*c <= x && x <= 105*c {
			return true
		}
	}
	return false
}

func main() {
	var x int
	fmt.Scan(&x)
	if solution(x) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
