package main

import (
	"fmt"
)

var (
	a, b, c, d, e, k int
)

func main() {
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)
	fmt.Scanf("%d", &d)
	fmt.Scanf("%d", &e)
	fmt.Scanf("%d", &k)

	if IsCommunicable(a, b, k) && IsCommunicable(a, c, k) && IsCommunicable(a, d, k) && IsCommunicable(a, e, k) &&
		IsCommunicable(b, c, k) && IsCommunicable(b, d, k) && IsCommunicable(b, e, k) &&
		IsCommunicable(c, d, k) && IsCommunicable(c, e, k) {
		fmt.Println("Yay!")
	} else {
		fmt.Println(":(")
	}
}

func IsCommunicable(i, j, k int) bool {
	if k < j-i {
		return false
	}
	return true
}
