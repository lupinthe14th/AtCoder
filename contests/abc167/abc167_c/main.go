package main

import (
	"fmt"
)

// bit 全探索
// 全探索にはDFSもあるが、部分集合を全探索する場合に有用なbit全探索を用いた
// 買う: 1
// 買わない: 0
//
// 0,1,2
// -----
// 0,0,0
// 0,0,1
// 0,1,0
// 0,1,1
// 1,0,0
// 1,0,1
// 1,1,0
// 1,1,1
// https://youtu.be/ENSOy8u9K9I?t=1302
func solution(n, m, x int, a [][]int, c []int) int {
	ans := 1 << 31
	// s: 部分集合
	// 1<<uint(n)
	// n = 3の場合, 1000 -> 8
	// 1<<n -> 2^n
	for s := 0; s < 1<<uint(n); s++ {
		cost := 0
		d := make([]int, m)
		for i := 0; i < n; i++ {
			// iビット目が1か調べる
			// 1. 右ビットシフト1してiビット目を一番右に移動する（最下位ビットにシフトする）
			// 2. 最下位ビットを取り出すため &1する（両方が1なら1）
			if s>>uint(i)&1 == 1 {
				cost += c[i]
				for j := 0; j < m; j++ {
					d[j] += a[i][j]
				}
			}
		}
		ok := true
		for i := 0; i < m; i++ {
			if d[i] < x {
				ok = false
			}
		}
		if ok {
			ans = min(ans, cost)
		}
	}
	// 答えが更新されていない場合は条件を満たす場合が無いので-1を返す
	if ans == 1<<31 {
		return -1
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
	var n, m, x int
	fmt.Scan(&n, &m, &x)

	c := make([]int, 0, n)
	a := make([][]int, n)

	for i := 0; i < n; i++ {
		tmp := 0
		fmt.Scan(&tmp)
		c = append(c, tmp)

		a[i] = make([]int, 0, m)
		for j := 0; j < m; j++ {
			tmp := 0
			fmt.Scan(&tmp)
			a[i] = append(a[i], tmp)
		}
	}
	fmt.Println(solution(n, m, x, a, c))
}
