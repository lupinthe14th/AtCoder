package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Parallel()
	type in struct {
		n, m int
		k, p []int
		s    [][]int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{n: 2, m: 2, k: []int{2, 1}, s: [][]int{{1, 2}, {2}}, p: []int{0, 1}}, want: 1},
		{in: in{n: 2, m: 3, k: []int{2, 1, 1}, s: [][]int{{1, 2}, {1}, {2}}, p: []int{0, 0, 1}}, want: 0},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.n, tt.in.m, tt.in.k, tt.in.p, tt.in.s)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}

func Example_main() {
	fd, _ := os.Open(filepath.Join("testdata", "input.txt"))
	orgStdin := os.Stdin
	os.Stdin = fd
	defer func() {
		os.Stdin = orgStdin
		fd.Close()
	}()

	main()

	// Output: 8
}
