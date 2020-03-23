package main

import (
	"fmt"
	"log"
)

func solution(n, m int) int {
	var ans int

	// n: 偶数
	even := make([]int, n)
	for i := range even {
		even[i] = 2
	}
	// m: 奇数
	odd := make([]int, m)
	for i := range odd {
		odd[i] = 1
	}

	log.Print(odd)
	log.Print(even)
	number := append(odd, even...)

	for i := 0; i < n+m; i++ {
		for j := 0; j < n+m; j++ {
			if i != j && (number[i]+number[j])&1 != 1 {
				ans++
			}
		}
	}
	return ans / 2
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	fmt.Println(solution(n, m))
}
