package main

import (
	"errors"
	"fmt"
	"strconv"
)

func solution(s string) (int, error) {
	var ans int
	n, err := strconv.Atoi(s)
	if err != nil {
		return ans, errors.New("error")
	}
	ans = n * 2
	return ans, nil
}

func main() {
	var s string
	fmt.Scan(&s)
	ans, err := solution(s)
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println(ans)
	}
}
