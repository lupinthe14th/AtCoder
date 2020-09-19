package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(s string) string {
	if s[len(s)-1] == 's' {
		return s + "es"
	}
	return s + "s"
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var s string
	fmt.Fscan(r, &s)
	fmt.Println(solution(s))
}
