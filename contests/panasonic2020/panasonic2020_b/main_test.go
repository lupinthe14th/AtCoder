package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type in struct {
	h, w int
}

type Case struct {
	in   in
	want int
}

var cases = []Case{
	{in: in{h: 4, w: 5}, want: 10},
	{in: in{h: 7, w: 3}, want: 11},
}

func TestSolution(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in.h, tt.in.w)
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

	// Output: 500000000000000000
}
