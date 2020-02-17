package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Refarence:
// - https://atcoder.jp/contests/abc155/submissions/10158108
// - https://atcoder.jp/contests/abc155/submissions/10150710
// - https://www.youtube.com/watch?v=SG60Cp9pSog&t=3175s
func solution(n, k int, nums []int) int64 {
	var l, r int64 = -1 << 63, 1<<63 - 1
	sort.Ints(nums)
	for l+1 < r {
		x := (l + r) >> 1
		var tot int64 = 0
		for i := 0; i < n; i++ {
			if nums[i] < 0 {
				l, r := -1, n
				for l+1 < r {
					c := (l + r) >> 1
					if int64(nums[c])*int64(nums[i]) < x {
						r = c
					} else {
						l = c
					}
				}
				tot += int64(n) - int64(r)
			} else {
				l, r := -1, n
				for l+1 < r {
					c := (l + r) >> 1
					if int64(nums[c])*int64(nums[i]) < x {
						l = c
					} else {
						r = c
					}
				}
				tot += int64(r)
			}
			if int64(nums[i])*int64(nums[i]) < x {
				tot--
			}
		}
		tot /= 2
		if tot < int64(k) {
			l = x
		} else {
			r = x
		}
	}
	return l
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	reader := bufio.NewReaderSize(os.Stdin, 2*1e6)
	str, _, _ := reader.ReadLine()
	nums := make([]int, n)
	for i, s := range strings.Split(string(str), " ") {
		nums[i], _ = strconv.Atoi(s)
	}

	fmt.Println(solution(n, k, nums))
}
