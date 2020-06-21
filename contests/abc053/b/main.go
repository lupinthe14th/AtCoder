package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(s string) int {
	l := len(s)
	f, r := 0, 0
	for i := 0; i < l; i++ {
		if s[i] == 'A' {
			f = i
			break
		}
	}

	for i := l - 1; i >= 0; i-- {
		if s[i] == 'Z' {
			r = i
			break
		}
	}
	return r - f + 1
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var s string
	fmt.Fscan(r, &s)
	fmt.Println(solution(s))
}
