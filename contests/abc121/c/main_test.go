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
		n, m   int
		matrix [][2]int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{n: 2, m: 5, matrix: [][2]int{{4, 9}, {2, 4}}}, want: 12},
		{in: in{n: 4, m: 30, matrix: [][2]int{{6, 18}, {2, 5}, {3, 10}, {7, 9}}}, want: 130},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.n, tt.in.m, tt.in.matrix)
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

	// Output: 100000000000000
}
