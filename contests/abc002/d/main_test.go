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
		xy   [][2]int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{n: 5, m: 3, xy: [][2]int{{1, 2}, {2, 3}, {1, 3}}}, want: 3},
		{in: in{n: 5, m: 3, xy: [][2]int{{1, 2}, {2, 3}, {3, 4}}}, want: 2},
		{in: in{n: 12, m: 0, xy: [][2]int{}}, want: 1},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.n, tt.in.m, tt.in.xy)
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

	// Output: 4
}
