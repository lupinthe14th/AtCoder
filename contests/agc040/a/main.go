package main

import (
	"bufio"
	"fmt"
	"os"
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
	var fp *os.File = os.Stdin
	reader := bufio.NewReaderSize(fp, 1024)
	s := readLine(reader)
	fmt.Println(solution(s))
}

func readLine(reader *bufio.Reader) string {
	buf := make([]byte, 0, 1024)

	for {
		l, p, _ := reader.ReadLine()

		buf = append(buf, l...)

		if !p {
			break
		}
	}
	return string(buf)
}
