package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solution(n int, nums []int) int {
	var ans int = 1 << 31
	for i := 1; i <= 100; i++ {
		sum := 0
		for j := 0; j < n; j++ {
			x := nums[j] - i
			sum += x * x
		}
		ans = min(ans, sum)
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		nums[i], _ = strconv.Atoi(sc.Text())
	}

	fmt.Println(solution(n, nums))
}
