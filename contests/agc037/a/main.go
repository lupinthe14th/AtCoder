package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://atcoder.jp/contests/agc037/submissions/6957166
func solution(s string) int {
	out := 0
	b := []byte{}
	p := ""
	for i := 0; i < len(s); i++ {
		b = append(b, s[i])
		if string(b) == p {
			continue
		}
		p = string(b)
		b = []byte{}
		out++
	}
	return out
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var s string
	fmt.Fscan(r, &s)
	fmt.Println(solution(s))
}
