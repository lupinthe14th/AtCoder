package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func solution(s string) string {
	var stack units
	l := len(s)
	var b int
	for i := 1; i < l; i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			tmp := s[b : i+1]
			stack = append(stack, Unit{org: tmp, lower: low(tmp)})
			b = i + 1
			i = b
		}
	}
	sort.Sort(stack)
	ans := make([]byte, 0, l)
	for i := range stack {
		ans = append(ans, stack[i].org...)
	}
	return string(ans)
}

type Unit struct {
	org, lower string
}

type units []Unit

func (u units) Len() int           { return len(u) }
func (u units) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }
func (u units) Less(i, j int) bool { return u[i].lower < u[j].lower }

func low(s string) string {
	ans := make([]byte, len(s))
	for i := range s {
		ans[i] = byte(s[i] | ' ')
	}
	return string(ans)
}
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1e5)
	str, _, _ := reader.ReadLine()
	fmt.Println(solution(string(str)))
}
