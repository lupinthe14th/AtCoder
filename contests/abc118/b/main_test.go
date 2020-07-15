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
		k    []int
		a    [][]int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{n: 3, m: 4, k: []int{2, 3, 2}, a: [][]int{{1, 3}, {1, 2, 3}, {3, 2}}}, want: 1},
		{in: in{n: 1, m: 30, k: []int{3}, a: [][]int{{5, 10, 30}}}, want: 3},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.n, tt.in.m, tt.in.k, tt.in.a)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}

func Example_main() {
	fd, _ := os.Open("./input.txt")
	orgStdin := os.Stdin
	os.Stdin = fd
	defer func() {
		os.Stdin = orgStdin
		fd.Close()
	}()

	main()

	// Output: 0
}
