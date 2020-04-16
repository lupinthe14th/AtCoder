package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type in struct {
	n, a, b int64
}

type Case struct {
	in   in
	want int64
}

var cases = []Case{
	{in: in{n: 8, a: 3, b: 4}, want: 4},
	{in: in{n: 8, a: 0, b: 4}, want: 0},
	{in: in{n: 6, a: 2, b: 4}, want: 2},
	{in: in{n: 6, a: 2, b: 0}, want: 6},
}

func TestSolution(t *testing.T) {
	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(c.in.n, c.in.a, c.in.b)
			if got != c.want {
				t.Errorf("in: %+v, got: %d, want: %d", c.in, got, c.want)
			}
		})
	}
}

func Example_main() {
	c, _ := ioutil.ReadFile("./input.txt")
	inr, inw, _ := os.Pipe()
	orgStdin := os.Stdin
	inw.Write(c)
	inw.Close()
	os.Stdin = orgStdin
	defer func() { os.Stdin = orgStdin }()
	os.Stdin = inr

	main()

	// Output: 0
}
