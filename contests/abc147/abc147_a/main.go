package main

import (
	"fmt"
)

func blackjack(one, two, three int) string {
	sum := one + two + three
	if sum >= 22 {
		return "bust"
	}
	return "win"
}

func main() {
	var one, two, three int
	fmt.Scan(&one, &two, &three)
	fmt.Println(blackjack(one, two, three))
}
