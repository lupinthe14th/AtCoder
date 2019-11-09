package main

import (
	"fmt"
)

func sumOfTwoIntegers(sum uint) uint {
	return (sum - 1) >> 1
}

func main() {
	var n uint
	fmt.Scan(&n)
	fmt.Println(sumOfTwoIntegers(n))
}
