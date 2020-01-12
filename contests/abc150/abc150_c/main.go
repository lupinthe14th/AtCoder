package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
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
	var i int
	for len(nums) != 0 {
		n := nums[0]
		log.Printf("n: %d", n)
		a := ascNumber(nums, n) - 1
		if a == 0 {
			ans += 1
		} else {
			ans += a * factorical(l-i-1)
		}
		log.Printf("i: %d, ans: %d", i, ans)
		nums = nums[1:]
		i++
	}
	log.Printf("ans: %d", ans+1)
	return ans + 1
}

func ascNumber(nums []int, n int) int {
	sort.Ints(nums)
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
