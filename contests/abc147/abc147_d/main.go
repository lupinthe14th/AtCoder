package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const MOD = 1e9 + 7

func xorSum4(n int, a []int) int {

	var ans int
	for d := 0; d < 61; d++ {
		var zero, one int
		x := 1 << uint(d)
		for i := range a {
			if 0 < (a[i] & x) {
				one++
			} else {
				zero++
			}
		}
		// Devide because it overflows.
		// SeeAlso: https://qiita.com/drken/items/3b4fdf0a78e7a138cd9a#オーバーフローに注意
		v := (zero * one % MOD) * (x % MOD) % MOD
		ans = (ans + v) % MOD
	}
	return ans
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 10240*10240)

	n, err := strconv.Atoi(readLine(reader))
	checkError(err)
	a := make([]int, n)
	ss := strings.Split(readLine(reader), " ")
	for i := range ss {
		a[i], err = strconv.Atoi(ss[i])
		checkError(err)
	}
	fmt.Println(xorSum4(n, a))
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
