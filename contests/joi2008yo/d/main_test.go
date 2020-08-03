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
		m, n int
		p, q [][2]int
	}
	tests := []struct {
		in   in
		want [2]int
	}{
		{in: in{m: 5, n: 10, p: [][2]int{{8, 5}, {6, 4}, {4, 3}, {7, 10}, {0, 10}}, q: [][2]int{{10, 5}, {2, 7}, {9, 7}, {8, 10}, {10, 2}, {1, 2}, {8, 1}, {6, 7}, {6, 0}, {0, 9}}}, want: [2]int{2, -3}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.m, tt.in.n, tt.in.p, tt.in.q)
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

	// Output: -384281 179674
}
