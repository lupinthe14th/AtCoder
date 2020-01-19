package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func solution(n int, p []int) int {
	var c int
	var low int = 1<<31 - 1
	for i := 1; i <= n; i++ {
		pi := p[i-1]
		if low > pi {
			c++
			low = pi
		}
	}
	return c
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 10240*10240)

	n, err := strconv.Atoi(readLine(reader))
	checkError(err)
	p := make([]int, n)
	ss := strings.Split(readLine(reader), " ")
	for i := range ss {
		p[i], err = strconv.Atoi(ss[i])
		checkError(err)
	}
	fmt.Println(solution(n, p))
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
