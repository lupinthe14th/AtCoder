package main

import (
	"fmt"
)

func solution(k, n int, nums []int) int {
	memo := make([]int, n)
	memo[0] = nums[0] + k - nums[len(nums)-1]

	for i := 1; i < n; i++ {
		memo[i] = nums[i] - nums[i-1]
	}
	maxSoFar := memo[0]

	for i := 1; i < n; i++ {
		if memo[i] > maxSoFar {
			maxSoFar = memo[i]
		}
	}
	return k - maxSoFar
}

func main() {
	var k, n int
	fmt.Scan(&k, &n)
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		t := 0
		fmt.Scan(&t)
		nums = append(nums, t)
	}

	fmt.Println(solution(k, n, nums))
}
