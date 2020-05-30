package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(s string) int {
	c := 0
	out := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'B' {
			c++
		} else {
			out += c
		}
	}
	return out
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
