package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solution(n, k int, h []int) int {
	if n <= k {
		return 0
	}
	sort.Ints(h)
	return sliceSum(h[:n-k])
}

func sliceSum(s []int) int {
	var sum int
	for _, v := range s {
		sum += v
	}
	return sum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 10240*10240)

	var n, k int
	var err error
	fmt.Scan(&n, &k)
	h := make([]int, n)
	ss := strings.Split(readLine(reader), " ")
	for i := range ss {
		h[i], err = strconv.Atoi(ss[i])
		checkError(err)
	}
	fmt.Println(solution(n, k, h))
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
