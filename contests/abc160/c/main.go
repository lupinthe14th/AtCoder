package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solution(k, n int, nums []int) int {
	var ans int = 1 << 31
	var sum int
	for i := 0; i < n; i++ {
		for j := i; j < n-1; j++ {
			if j+1 == n {
				sum += k - nums[j] + nums[0]
			} else {
				sum += nums[j+1] - nums[j]
			}
		}
		ans = min(ans, sum)
	}
	for i := 0; i < n; i++ {
		for j := i; j < n-1; j-- {
			if j == 0 {
				sum += k - nums[j] - nums[n-1]
			} else {
				sum += nums[j] - nums[j-1]
			}
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
	k, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		nums[i], _ = strconv.Atoi(sc.Text())
	}

	fmt.Println(solution(k, n, nums))
}
