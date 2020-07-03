package main

import (
	"fmt"
)

func solution(t int, matrix [][]int) []int64 {
	out := make([]int64, 0)
	memo := make(map[int64]int64, 0)
	var a, b, c, d int64
	var helper func(n int64) int64
	helper = func(n int64) int64 {

		if n == 0 {
			return 0
		}
		if n == 1 {
			return d
		}

		if k, ok := memo[n]; ok {
			return k
		}

		var l2, r2, l3, r3, l5, r5 int64

		l2 = (n / 2) * 2
		r2 = ((n + 1) / 2) * 2
		l3 = (n / 3) * 3
		r3 = ((n + 2) / 3) * 3
		l5 = (n / 5) * 5
		r5 = ((n + 4) / 5) * 5

		var res int64 = 1e18
		if n < res/d {
			res = n * d
		}

		res = min(res, abs(l2-n)*d+a+helper(l2/2))
		res = min(res, abs(r2-n)*d+a+helper(r2/2))

		res = min(res, abs(l3-n)*d+b+helper(l3/3))
		res = min(res, abs(r3-n)*d+b+helper(r3/3))

		res = min(res, abs(l5-n)*d+c+helper(l5/5))
		res = min(res, abs(r5-n)*d+c+helper(r5/5))

		memo[n] = res
		return res
	}
	for i := range matrix {
		memo = make(map[int64]int64, 0)
		n := int64(matrix[i][0])
		a = int64(matrix[i][1])
		b = int64(matrix[i][2])
		c = int64(matrix[i][3])
		d = int64(matrix[i][4])
		out = append(out, helper(n))
	}

	return out
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func main() {
	var t int
	fmt.Scan(&t)
	matrix := make([][]int, t)
	for i := range matrix {
		var n, a, b, c, d int
		fmt.Scan(&n, &a, &b, &c, &d)
		matrix[i] = make([]int, 5)
		matrix[i] = []int{n, a, b, c, d}
	}
	for _, a := range solution(t, matrix) {
		fmt.Println(a)
	}
}
