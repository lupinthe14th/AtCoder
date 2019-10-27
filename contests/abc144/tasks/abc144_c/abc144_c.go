package main

import (
	"fmt"
)

func walkOnMultiplicationTable(num int) int {

	l := num
	x := 1
	for i := 1; i < l; i++ {
		if num%i == 0 {
			x = i
		}
		l = num / i
	}
	return x + num/x - 2

}

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println(walkOnMultiplicationTable(a))
}
