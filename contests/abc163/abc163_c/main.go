package main

import (
	"fmt"
)

func solution(n int, nums []int) []int {
	nums = append([]int{0}, nums...)
	bosses := make([]int, n)
	m := make(map[int]int)
	for _, num := range nums {
		m[num-1]++
	}
	for i := 0; i < n; i++ {
		bosses[i] = m[i]
	}
	return bosses
}

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		a := 0
		fmt.Scan(&a)
		nums = append(nums, a)
	}
	for _, num := range solution(n, nums) {
		fmt.Println(num)
	}
}
