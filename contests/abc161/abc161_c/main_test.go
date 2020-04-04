package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestSolution(t *testing.T) {
	type in struct {
		n, k int
	}

	var cases = []struct {
		in   in
		want int
	}{
		{in: in{n: 7, k: 4}, want: 1},
		{in: in{n: 2, k: 6}, want: 2},
		{in: in{n: 101, k: 6}, want: 1},
		{in: in{n: 102, k: 6}, want: 0},
		{in: in{n: 102, k: 7}, want: 3},
		{in: in{n: 100, k: 1}, want: 0},
		{in: in{n: 100, k: 2}, want: 0},
		{in: in{n: 3, k: 100}, want: 3},
		{in: in{n: 0, k: 100}, want: 0},
		{in: in{n: 0, k: 1}, want: 0},
		{in: in{n: 1, k: 1}, want: 0},
	}
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in.n, tt.in.k)
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

	// Output: 0
}
