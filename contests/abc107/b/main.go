package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(h, w int, a []string) [][]byte {
	memo := make([][]int, h+1)
	for i := range memo {
		memo[i] = make([]int, w+1)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if a[i][j] == '.' {
				memo[i][j]++
			}
		}
	}

	for i := 0; i < h+1; i++ {
		sum := 0
		for j := 0; j < w+1; j++ {
			sum += memo[i][j]
			if j == w {
				memo[i][j] = sum
			}
		}
	}
	for j := 0; j < w+1; j++ {
		sum := 0
		for i := 0; i < h+1; i++ {
			sum += memo[i][j]
			if i == h {
				memo[i][j] = sum
			}
		}
	}

	b := make([][]byte, h)
	k := 0
	for i := 0; i < h; i++ {
		if memo[i][w] != w {
			b[k] = []byte{}
			k++
		}
		for j := 0; j < w; j++ {
			if memo[i][w] != w && memo[h][j] != h {
				b[k-1] = append(b[k-1], a[i][j])
			}
		}
	}

	if k != h {
		b = b[:k]
	}
	return b
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var h, w int
	fmt.Fscan(r, &h, &w)
	a := make([]string, h)
	for i := range a {
		fmt.Fscan(r, &a[i])
	}
	for _, b := range solution(h, w, a) {
		fmt.Println(string(b))
	}
}
