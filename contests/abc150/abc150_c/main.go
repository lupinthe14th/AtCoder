package main

import (
	"fmt"
	"sort"
)

func solution(n int, p, q []int) int {
	i := dictionaryOrder(p)
	j := dictionaryOrder(q)
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
	for i := 0; i < l; i++ {
		n := nums[i]
		arr := make([]int, len(nums)-i)
		copy(arr, nums[i:])
		a := ascNumber(arr, n) - 1
		if a == 0 {
			ans += 0
		} else {
			ans += a * factorical(l-i-1)
		}
	}
	ans++
	return ans
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

	var n int
	fmt.Scan(&n)
	p := make([]int, n)
	for i := range p {
		fmt.Scan(&p[i])
	}
	q := make([]int, n)
	for i := range q {
		fmt.Scan(&q[i])
	}
	fmt.Println(solution(n, p, q))
}
