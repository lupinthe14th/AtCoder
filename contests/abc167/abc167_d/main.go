package main

import (
	"fmt"
)

// https://youtu.be/ENSOy8u9K9I?t=2553
// other solution: ダブリング: https://youtu.be/ENSOy8u9K9I?t=3600
func solution(n, k int, a []int) int {
	// 訪れた頂点をメモ
	visit := make([]int, 0)
	// 訪れた頂点のホップ数をメモ
	hop := make([]int, n+1)
	for i := range hop {
		// 初期値として訪れていない印をメモ
		hop[i] = -1
	}

	// cur: 現在地点。初期値は1
	// c: 周期
	// 例外部分の長さ
	cur := 1
	c := 1
	l := 0
	for hop[cur] == -1 {
		// ホップ数を記録。この時の長さを使うと簡単（カウンターいらない）
		hop[cur] = len(visit)
		// 訪れた頂点を記録
		visit = append(visit, cur)
		//ワープ
		cur = a[cur-1]
	}

	c = len(visit) - hop[cur]
	l = hop[cur]
	ans := 0
	if k < l {
		ans = visit[k]
	} else {
		k -= l
		k %= c
		ans = visit[l+k]
	}
	return ans
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	a := make([]int, 0, n)
	for i := 0; i < n; i++ {
		tmp := 0
		fmt.Scan(&tmp)
		a = append(a, tmp)
	}
	fmt.Println(solution(n, k, a))
}
