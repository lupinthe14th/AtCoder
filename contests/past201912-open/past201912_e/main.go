package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, q int
	fmt.Scan(&n, &q)

	f := make([][]string, n)
	for i := range f {
		f[i] = make([]string, n)
	}

	for i := range f {
		for j := range f[i] {
			f[i][j] = "N"
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")

		switch {
		case s[0] == "1":
			a, _ := strconv.Atoi(s[1])
			b, _ := strconv.Atoi(s[2])
			f[a-1][b-1] = "Y"
		case s[0] == "2":
			log.Printf("s: %s, a: %s", s[0], s[1])
			a, _ := strconv.Atoi(s[1])
			for _, v := range followerList(a-1, f) {
				if a-1 != v {
					f[a-1][v] = "Y"
				}
			}
		case s[0] == "3":
			log.Printf("s: %s, a: %s", s[0], s[1])
			a, _ := strconv.Atoi(s[1])
			for _, v := range followingList(a-1, f) {
				for _, vv := range followingList(v, f) {
					if a-1 != vv {
						f[a-1][vv] = "Y"
					}
				}
			}
		}
	}

	for i := range f {
		fmt.Println(strings.Join(f[i], ""))
	}
}

// ユーザーaをフォローしているユーザーのリストを返す
func followerList(a int, matrix [][]string) []int {
	ans := make([]int, 0)
	for i := range matrix {
		if i == a {
			continue
		}
		if matrix[i][a] == "Y" {
			ans = append(ans, i)
		}
	}
	return ans
}

// ユーザーaがフォローしているユーザーのリストを返す
func followingList(a int, matrix [][]string) []int {
	ans := make([]int, 0)
	for i := range matrix {
		if i == a {
			continue
		}
		if matrix[a][i] == "Y" {
			ans = append(ans, i)
		}
	}
	return ans
}
