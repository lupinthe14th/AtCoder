package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(s string) string {
	// N: 0
	// W: 1
	// S: 2
	// E: 3

	var d [4]int

	for i := range s {
		switch s[i] {
		case 'N':
			d[0]++
		case 'W':
			d[1]++
		case 'S':
			d[2]++
		case 'E':
			d[3]++
		}
	}
	if d[0] > 0 && d[2] == 0 || d[0] == 0 && d[2] > 0 {
		return "No"
	}
	if d[1] > 0 && d[3] == 0 || d[1] == 0 && d[3] > 0 {
		return "No"
	}
	return "Yes"
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var s string
	fmt.Fscan(r, &s)
	fmt.Println(solution(s))
}
