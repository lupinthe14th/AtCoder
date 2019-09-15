package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

func isTapDanceEasy(s string) bool {
	if len(s) > 1 && len(s) < 100 {
		odds := make([]string, 0)
		evens := make([]string, 0)
		sl := strings.Split(s, "")

		for i, char := range sl {
			if (i+1)%2 == 0 {
				evens = append(evens, char)
				continue
			}
			odds = append(odds, char)
		}
		odd := strings.Join(odds, "")
		log.Printf("odd: %s", odd)
		even := strings.Join(evens, "")
		log.Printf("even: %s", even)
		if strings.Contains(odd, "L") || strings.Contains(even, "R") {
			return false
		}
		return true
	}
	panic(errors.New("No solution"))
}

func main() {
	var s string
	fmt.Scan(&s)
	if isTapDanceEasy(s) {
		fmt.Println("Yes")
	}
	fmt.Println("No")
}
