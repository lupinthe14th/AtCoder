package main

import (
	"fmt"
)

func solution(n int) string {
	out := []byte{}
	for n > 0 {
		y := n % 26
		if y != 0 {
			out = append([]byte{byte(y-1) + 'a'}, out...)
			n /= 26
		} else {
			out = append([]byte{'z'}, out...)
			n /= 26
			n--
		}
	}

	return string(out)
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(solution(n))
}
