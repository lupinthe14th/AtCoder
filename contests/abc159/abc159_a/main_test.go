package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type in struct {
	n, m int
}

type Case struct {
	in   in
	want int
}

var cases = []Case{
	{in: in{n: 2, m: 1}, want: 1},
	{in: in{n: 4, m: 3}, want: 9},
	{in: in{n: 1, m: 1}, want: 0},
	{in: in{n: 13, m: 3}, want: 81},
	{in: in{n: 0, m: 3}, want: 3},
}

func TestSolution(t *testing.T) {
	t.Parallel()
	for i, tt := range cases {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.n, tt.in.m)
			if got != tt.want {
				t.Errorf("in: %+v, got: %d, want: %d", tt.in, got, tt.want)
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

	// Output: 9
}
