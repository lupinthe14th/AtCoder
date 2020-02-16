package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solution(n int, texts []string) []string {
	ans := make([]string, 0)
	var c int
	seen := make(map[string]int, n)

	for _, text := range texts {
		seen[text]++
		c = max(c, seen[text])
	}

	for k, v := range seen {
		if v == c {
			ans = append(ans, k)
		}
	}
	sort.Strings(ans)
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	texts := make([]string, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		texts[i] = sc.Text()
	}

	ans := solution(n, texts)
	for _, text := range ans {
		fmt.Println(text)
	}
}
