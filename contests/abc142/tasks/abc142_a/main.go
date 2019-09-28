package main

import (
	"errors"
	"fmt"
)

func oddPercent(num int) float64 {
	var odds = make([]int, 0, num)
	if num >= 1 && num <= 100 {
		for i := num; i >= 0; i-- {
			if i%2 == 1 {
				odds = append(odds, i)
			}
		}
		p := float64(len(odds)) / float64(num)
		return p
	}
	panic(errors.New("No solution"))
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Print(oddPercent(n))
}
