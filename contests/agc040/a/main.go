package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solution(s string) int {
	n := len(s) + 1
	a := make([]int, n)
	for i := 0; i < n-1; i++ {
		if s[i] == '<' {
			a[i+1] = max(a[i+1], a[i]+1)
		}
	}

	for i := n - 2; i >= 0; i-- {
		if s[i] == '>' {
			a[i] = max(a[i], a[i+1]+1)
		}
	}
	out := 0
	for i := 0; i < n; i++ {
		out += a[i]
	}
	return out
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 10240*10240)
	str, _, _ := reader.ReadLine()
	fmt.Println(solution(strings.TrimRight(string(str), "\r\n")))
}
