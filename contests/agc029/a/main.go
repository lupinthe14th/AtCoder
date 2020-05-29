package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	reader := bufio.NewReaderSize(os.Stdin, 10240*10240)
	str, _, _ := reader.ReadLine()
	fmt.Println(solution(strings.TrimRight(string(str), "\r\n")))
}
