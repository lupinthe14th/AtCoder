package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func rollerCoaster(n, k int, l []int) int {
	stack := make([]int, 0, n)
	for _, v := range l {
		if v >= k {
			stack = append(stack, v)
		}
	}
	return len(stack)
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	high := strings.Split(readLine(reader), " ")

	var h = make([]int, 0)
	for _, v := range high {
		num, err := strconv.ParseInt(v, 10, 64)
		checkError(err)
		h = append(h, int(num))
	}
	fmt.Print(rollerCoaster(n, k, h))
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
