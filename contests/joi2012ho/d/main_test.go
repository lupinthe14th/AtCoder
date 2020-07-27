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
		g    [][3]int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{n: 5, m: 2, g: [][3]int{{2, 2, 1}, {2, 1, 3}}}, want: 12},
		{in: in{n: 5, m: 3, g: [][3]int{{1, 1, 2}, {2, 1, 2}, {4, 2, 1}}}, want: 11},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.n, tt.in.m, tt.in.g)
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

	// Output: 12
}
