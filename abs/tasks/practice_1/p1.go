package main

import "fmt"

var (
	a, b, c int
	s       string
)

func printSumAndWord(a, b, c int, s string) string {
	return fmt.Sprintf("%d %s", a+b+c, s)
}

func main() {
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d %d", &b, &c)
	fmt.Scanf("%s", &s)
	fmt.Println(printSumAndWord(a, b, c, s))
}
