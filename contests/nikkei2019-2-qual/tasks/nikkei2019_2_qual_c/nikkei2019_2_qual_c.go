package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Perm struct {
	idx int
	num int
}

type Perms []Perm

func (p Perms) Len() int {
	return len(p)
}

func (p Perms) Less(i, j int) bool {
	return p[i].num < p[j].num
}

func (p Perms) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func swaps(n int, a, b []int) bool {
	// まず巡回置換 (p1, ..., pN ) を Ap1 ≦ Ap2 ≦ ... ≦ ApN となるようにとる
	ap := make([]int, 0, n)
	for _, v := range a {
		ap = append(ap, v)
	}
	sort.Ints(ap)
	bp := make([]int, 0, n)
	for _, v := range b {
		bp = append(bp, v)
	}
	sort.Ints(bp)

	// このとき Api ≦ Bi(i = 1, ..., N ) が成り立たなければ、1 つめの条件をみたす (p1 , ..., pN ) が存在しないので、false
	for i := 0; i < n; i++ {
		if ap[i] > bp[i] {
			return false
		}

	}

	// Api ≦ Bi(i = 1,...,N) となる場合
	// 巡回置換(p1,...,pN)をサイクルに分解したときに、2つ以上のサイクルに分解されるならば、true
	// N-1回あればできるやつ
	// 目標位置をswapできる場所があれば、達成可能
	for i := 0; i < n-1; i++ {
		if ap[i+1] <= bp[i] {
			return true
		}
	}

	// この時点で、Aは相異なり、目標インデックスも1つに決まる
	// a をsortすると順序が変わってしまうので最初に処理する
	// 時間計算量が増える。。。
	perms := make([]Perm, n)
	for i, v := range a {
		perms[i] = Perm{idx: i, num: v}
	}

	sort.Sort(Perms(perms))

	p := make([]int, 0, n)
	for _, perm := range perms {
		p = append(p, perm.idx)
	}
	count := 0
	x := 0
	for x != 0 {
		count++
		x = p[x]
	}

	if count <= n-2 {
		return true
	}
	return false
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var n int
	if sc.Scan() {
		n, _ = strconv.Atoi(sc.Text())
	}

	a := make([]int, n)
	if sc.Scan() {
		ss := strings.Split(sc.Text(), " ")
		for i := range ss {
			a[i], _ = strconv.Atoi(ss[i])
		}
	}

	b := make([]int, n)
	if sc.Scan() {
		ss := strings.Split(sc.Text(), " ")
		for i := range ss {
			b[i], _ = strconv.Atoi(ss[i])
		}
	}

	if swaps(n, a, b) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
