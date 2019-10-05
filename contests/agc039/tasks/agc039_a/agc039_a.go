package main

import (
	"fmt"
)

func connectionAndDisconnection(s string, k int) int {

	if len(s) == 1 {
		return k / 2
	}
	r := []rune(s)
	counter := 0
	j := 0
	for i := 1; i < len(r); i++ {
		if r[j] == r[i] {
			counter++
			if i+1 < len(r) && r[i+1] == r[i]+1 {
				r[i] = r[i] + 2
			} else {
				r[i] = r[i] + 1
			}
		}
		j++
	}
	return counter * k
}

func main() {
	var s string
	var k int
	fmt.Scan(&s)
	fmt.Scan(&k)
	fmt.Println(connectionAndDisconnection(s, k))
}
