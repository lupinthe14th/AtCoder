package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func solution(n int, p, q []int) int {
	i := dictionaryOrder(p)
	j := dictionaryOrder(q)
	log.Printf("i: %d, j: %d", i, j)
	return abs(i, j)
}

func abs(x, y int) int {
	z := x - y
	if z < 0 {
		return z * -1
	}
	return z
}

func dictionaryOrder(nums []int) int {
	var ans int
	l := len(nums)
	arr := make([]int, l)
	for i := 1; i <= l; i++ {
		arr[i-1] = i
	}
	log.Printf("arr: %v", arr)
	for i := 0; i < l; i++ {
		n := nums[i]
		log.Printf("n: %d", n)
		a := ascNumber(arr, n) - 1
		log.Printf("ascNumber: %d", a)
		log.Printf("nums: %v", nums)
		if a == 0 {
			ans += 0
		} else {
			ans += a * factorical(l-i-1)
		}
		log.Printf("l: %d, ans: %d", l, ans)
	}
	ans++
	log.Printf("ans: %d", ans)
	return ans
}

func ascNumber(nums []int, n int) int {
	var i int = 1
	for _, v := range nums {
		if v == n {
			break
		}
		i++
	}
	return i
}

func factorical(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorical(n-1)
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
	q := make([]int, n)
	ss = strings.Split(readLine(reader), " ")
	for i := range ss {
		q[i], err = strconv.Atoi(ss[i])
		checkError(err)
	}
	fmt.Println(solution(n, p, q))
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
