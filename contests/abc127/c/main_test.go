package main

import (
	"fmt"
	"os"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Parallel()
	type in struct {
		n, m int
		lr   [][]int
	}

	var cases = []struct {
		in   in
		want int
	}{
		{in: in{n: 4, m: 2, lr: [][]int{{1, 3}, {2, 4}}}, want: 2},
		{in: in{n: 10, m: 3, lr: [][]int{{3, 6}, {5, 7}, {6, 9}}}, want: 1},
		{in: in{n: 100000, m: 1, lr: [][]int{{1, 100000}}}, want: 100000},
	}
	for i, tt := range cases {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.n, tt.in.m, tt.in.lr)
			if got != tt.want {
				t.Errorf("in: %+v, got: %v, want: %v", tt.in, got, tt.want)
			}
		})
	}
}

func Example_main() {
	fd, _ := os.Open("./input.txt")
	orgStdin := os.Stdin
	os.Stdin = orgStdin
	os.Stdin = fd
	defer func() {
		os.Stdin = orgStdin
		fd.Close()
	}()

	main()

	// Output: 0
}
