package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type in struct {
	n, t int
	m    []wt
}

type Case struct {
	in   in
	want int
}

var cases = []Case{
	{in: in{n: 3, t: 7, m: []wt{{a: 2, b: 0, c: 2}, {a: 3, b: 2, c: 6}, {a: 0, b: 3, c: 3}}}, want: 2},
	{in: in{n: 1, t: 3, m: []wt{{a: 0, b: 3, c: 3}}}, want: 0},
}

func TestSolution(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in.n, tt.in.t, tt.in.m)
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

	// Output: 5
}
