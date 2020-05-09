package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestSolution(t *testing.T) {
	type in struct {
		a, b, c int
	}

	var cases = []struct {
		in   in
		want int
	}{
		{in: in{a: 4, b: 12, c: 20}, want: 3},
		{in: in{a: 1, b: 1, c: 1}, want: 0},
		{in: in{a: 454, b: 414, c: 444}, want: 1},
	}
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in.a, tt.in.b, tt.in.c)
			if got != tt.want {
				t.Errorf("in: %+v, got: %v, want: %v", tt.in, got, tt.want)
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

	// Output: -1
}
