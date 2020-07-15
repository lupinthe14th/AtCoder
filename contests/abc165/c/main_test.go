package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Parallel()
	type in struct {
		n, m, q int
		matrix  [][]int
	}
	var cases = []struct {
		in   in
		want int
	}{
		{in: in{n: 3, m: 4, q: 3, matrix: [][]int{{1, 3, 3, 100}, {1, 2, 2, 10}, {2, 3, 2, 10}}}, want: 110},
		{in: in{n: 10, m: 10, q: 1, matrix: [][]int{{1, 10, 9, 1}}}, want: 1},
	}
	for i, tt := range cases {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.n, tt.in.m, tt.in.q, tt.in.matrix)
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
	os.Stdin = orgStdin
	defer func() {
		inw.Close()
		os.Stdin = orgStdin
	}()
	os.Stdin = inr

	main()

	// Output: 357500
}
