package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func go2School(n int, l []int) []int {
	return []int{1, 2, 3}
}

func main() {
	var n int
	fmt.Scan(&n)

	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	human := strings.Split(readLine(reader), " ")

	var h = make([]int, 0)
	for _, v := range human {
		num, err := strconv.ParseInt(v, 10, 64)
		checkError(err)
		h = append(h, int(num))
	}
	for _, v := range go2School(n, h) {
		fmt.Printf("%d ", v)
	}
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
