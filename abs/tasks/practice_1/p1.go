package main

import "fmt"

var (
	a, b, c int
	s       string
)

func main() {
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d %d", &b, &c)
	fmt.Scanf("%s", &s)
	fmt.Printf("%d %s\n", a+b+c, s)
}
