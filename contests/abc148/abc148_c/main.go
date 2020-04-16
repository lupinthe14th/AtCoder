package main

import (
	"fmt"
)

func snack(a, b int) int {
	return (a * b) / gcd(a, b)
}

func gcd(a, b int) int {
	var tmp int
	for a%b != 0 {
		tmp = b
		b = a % b
		a = tmp
	}
	return b

}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(snack(a, b))
}
