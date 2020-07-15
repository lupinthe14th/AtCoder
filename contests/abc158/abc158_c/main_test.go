package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type in struct {
	a, b int
}

type Case struct {
	in   in
	want int
}

var cases = []Case{
	{in: in{a: 2, b: 2}, want: 25},
	{in: in{a: 19, b: 99}, want: -1},
}

func TestSolution(t *testing.T) {
	t.Parallel()
	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(c.in.a, c.in.b)
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

	// Output: 100
}
