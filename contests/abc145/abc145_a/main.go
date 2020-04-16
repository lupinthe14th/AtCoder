package main

import (
	"fmt"
)

func circle(r int) int {
	return r * r
}

func main() {
	var r int
	fmt.Scan(&r)
	fmt.Println(circle(r))
}
