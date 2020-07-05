package main

import (
	"fmt"
)

func solution(n int, s []string) []string {
	var ac, wa, tle, re int
	for _, t := range s {
		switch t {
		case "AC":
			ac++
		case "WA":
			wa++
		case "TLE":
			tle++
		case "RE":
			re++
		}
	}
	return []string{fmt.Sprintf("AC x %v", ac), fmt.Sprintf("WA x %v", wa), fmt.Sprintf("TLE x %v", tle), fmt.Sprintf("RE x %v", re)}
}

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]string, n)
	for i := range s {
		fmt.Scan(&s[i])
	}
	for _, t := range solution(n, s) {
		fmt.Println(t)
	}
}
