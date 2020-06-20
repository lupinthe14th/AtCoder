package main

import (
	"fmt"
)

func solution(o, e string) string {
	out := make([]byte, len(o)+len(e))
	for i := range out {
		if (i-1)&1 == 1 {
			out[i] = o[i/2]
		} else {
			out[i] = e[i/2]
		}
	}
	return string(out)
}

func main() {
	var o, e string
	fmt.Scan(&o, &e)
	fmt.Println(solution(o, e))
}
