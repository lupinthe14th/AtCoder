package main

import (
	"fmt"
)

var nums = []int{1, 1, 1, 2, 1, 2, 1, 5, 2, 2, 1, 5, 1, 2, 1, 14, 1, 5, 1, 5, 2, 2, 1, 15, 2, 2, 5, 4, 1, 4, 1, 51}

func solution(n int, nums []int) int {
	return nums[n-1]
}

func main() {
	var k int
	fmt.Scan(&k)
	fmt.Println(solution(k, nums))
}
