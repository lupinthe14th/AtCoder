package main

import (
	"bufio"
	"fmt"
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

	high := bufio.NewScanner(os.Stdin)
	high.Scan()
	var h = make([]int, 0)
	for _, v := range strings.Split(high.Text(), " ") {
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		h = append(h, n)
	}
	fmt.Print(rollerCoaster(n, k, h))
}
