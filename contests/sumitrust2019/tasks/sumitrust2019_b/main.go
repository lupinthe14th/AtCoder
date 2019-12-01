package main

import (
	"fmt"
	"math"
)

func taxRate(n int) (int, string) {
	tax := 1.08
	tmp := float64(n) / tax

	ans := math.Trunc(tmp)

	if int(ans*tax) == n {
		return int(ans), ""
	}
	if int((ans+1)*tax) == n {
		return int(ans + 1), ""
	}
	return 0, ":("
}

func main() {
	var n int
	fmt.Scan(&n)
	c, err := taxRate(n)
	if err != "" {
		fmt.Println(err)
	} else {
		fmt.Println(c)
	}
}
