package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type in struct {
	s, l, r int
}

type Case struct {
	in   in
	want int
}

var cases = []Case{
	{in: in{s: 5, l: 1, r: 5}, want: 5},
	{in: in{s: 2, l: 7, r: 10}, want: 7},
}

func TestSolution(t *testing.T) {
	t.Parallel()
	for i, tt := range cases {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.s, tt.in.l, tt.in.r)
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
