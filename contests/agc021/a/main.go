package main

import (
	"fmt"
)

func solution(n string) int {
	l := len(n)
	nums := make([]int, l)
	sum := 0
	for i := range n {
		nums[i] = int(n[i] - '0')
		sum += nums[i]
	}

	for i := l - 1; i > 0; i-- {
		nums[i] = 9
		nums[i-1] = nums[i-1] - 1
		tmp := 0
		for j := range nums {
			tmp += nums[j]
		}
		sum = max(sum, tmp)
	}
	return sum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var n string
	fmt.Scan(&n)
	fmt.Println(solution(n))
}
