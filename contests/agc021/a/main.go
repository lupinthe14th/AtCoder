package main

import (
	"fmt"
)

// SeeAlso: https://atcoder.jp/contests/agc021/submissions/2130303
func solution(n int) int {
	return max(helper(n), helper(dec(n)))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func helper(n int) int {
	out := 0
	for n > 0 {
		out += n % 10
		n /= 10
	}
	return out
}

func dec(n int) int {
	k := 1
	for n >= 10 {
		n /= 10
		k *= 10
	}
	return n*k - 1

}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(solution(n))
}
