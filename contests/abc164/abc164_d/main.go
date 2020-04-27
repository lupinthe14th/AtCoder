package main

import (
	"fmt"
)

// https://www.youtube.com/watch?v=Ra0y0EZ24ZM&t=460s
// https://atcoder.jp/contests/abc164/submissions/12362793
func solution(s string) int {
	l := len(s)
	sum_of_digits := 0
	cnts := make([]int, 2019)
	cnts[0] = 1

	d := 1
	for i := l - 1; i >= 0; i-- {
		sum_of_digits += (int(s[i]) - 48) * d
		sum_of_digits %= 2019
		d *= 10
		d %= 2019
		cnts[sum_of_digits] += 1
	}

	ans := 0
	for _, cnt := range cnts {
		ans += cnt * (cnt - 1) / 2
	}
	return ans
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(solution(s))
}
