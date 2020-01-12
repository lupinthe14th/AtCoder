package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func solution(n, k, m int, a []int) int {
	var ans int

	var ca, s int
	s = sum(a)
	ca = s / n

	if ca >= m {
		return 0
	} else {
		ans = m*n - s
	}
	if ans > k {
		ans = -1
	}
	return ans
}

func sum(a []int) int {
	var sum int
	for _, v := range a {
		sum += v
	}
	return sum
}

func main() {
	var n, k, m int
	fmt.Scan(&n, &k, &m)
	reader := bufio.NewReaderSize(os.Stdin, 10240*10240)
	a := make([]int, n)
	var err error
	ss := strings.Split(readLine(reader), " ")
	for i := range ss {
		a[i], err = strconv.Atoi(ss[i])
		checkError(err)
	}
	fmt.Println(solution(n, k, m, a))
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
