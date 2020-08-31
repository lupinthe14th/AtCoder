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
		ab   [][2]int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{n: 5, m: 3, ab: [][2]int{{1, 2}, {3, 4}, {5, 1}}}, want: 3},
		{in: in{n: 10, m: 4, ab: [][2]int{{3, 1}, {4, 1}, {5, 9}, {2, 6}}}, want: 3},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.n, tt.in.m, tt.in.ab)
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
