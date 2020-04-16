package main

import (
	"fmt"
	"reflect"
	"strings"
)

func echo(n int, s string) bool {
	if n == 1 {
		return false
	}
	sa := strings.Split(s, "")

	if reflect.DeepEqual(sa[:n/2], sa[n/2:]) {
		return true
	}

	return false
}

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)
	if echo(n, s) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
