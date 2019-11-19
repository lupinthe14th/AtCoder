// SeeAlso:
// - https://img.atcoder.jp/abc145/editorial.pdf
// - http://drken1215.hatenablog.com/entry/2018/06/08/210000
package main

import (
	"fmt"
)

const MAX = 1000000
const MOD = 1000000007

var fac = make([]int, MAX)
var finv = make([]int, MAX)
var inv = make([]int, MAX)

func knight(x, y int) int {
	// 1回の移動でX座標+Y座標の値は3増える。
	// よって、　X＋Yが3の倍数でない時答えは0
	if (x+y)%3 != 0 {
		return 0
	}
	// 3の倍数のとき、
	//(1,2)の移動の回数をn
	//(2,1)の移動の回数をmとすると
	// n + 2m = X
	// 2n + m = Y
	// が得られる
	m := (2*x - y) / 3
	n := x - 2*m
	return com(n+m, m)
}

// テーブルを作る前処理
func cominit() {
	fac[0] = 1
	fac[1] = 1

	finv[0] = 1
	finv[1] = 1
	inv[1] = 1
	for i := 2; i < MAX; i++ {
		fac[i] = fac[i-1] * i % MOD
		inv[i] = MOD - inv[MOD%i]*(MOD/i)%MOD
		finv[i] = finv[i-1] * inv[i] % MOD
	}
}

// 二項係数計算
func com(n, k int) int {
	if n < k {
		return 0
	}
	if n < 0 || k < 0 {
		return 0
	}
	cominit()
	return fac[n] * (finv[k] * finv[n-k] % MOD) % MOD
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	fmt.Println(knight(x, y))
}
