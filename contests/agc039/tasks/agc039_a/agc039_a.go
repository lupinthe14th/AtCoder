package main

import (
	"fmt"
	"strings"
)

func connectionAndDisconnection(s string, k int) int {

	// all the characters in S are the same pattern
	m := make(map[rune]int, 26)

	for _, c := range s {
		m[c]++
	}
	if len(m) == 1 {
		return (len(s) * k) / 2
	}

	// If these characters are the same pattern
	r := []rune(s)
	cost := 0
	if s[0] == s[len(s)-1] {
		for i := 1; i < len(r); i++ {
			if r[i-1] == r[i] {
				cost++
				r[i] = '#'
			}
		}
		l := string(s[0])
		head := len(s) - len(strings.TrimLeft(s, l))
		tail := len(s) - len(strings.TrimRight(s, l))
		return cost*k - (head/2+tail/2-(head+tail)/2)*(k-1)

	}
	// If the first and last characters of S are different pattern.
	for i := 1; i < len(r); i++ {
		if r[i-1] == r[i] {
			cost++
			r[i] = '#'
		}
	}
	return cost * k
}

func main() {
	var s string
	var k int
	fmt.Scan(&s)
	fmt.Scan(&k)
	fmt.Println(connectionAndDisconnection(s, k))
}
