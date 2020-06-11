package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestSolution(t *testing.T) {
	type in struct {
		n, d int
		x    [][]int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{n: 3, d: 2, x: [][]int{{1, 2}, {5, 5}, {-2, 8}}}, want: 1},
		{in: in{n: 5, d: 1, x: [][]int{{1}, {2}, {3}, {4}, {5}}}, want: 10},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in.n, tt.in.d, tt.in.x)
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

	// Output: 2
}
