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
		m1, d1, m2, d2 int
	}

	var tests = []struct {
		in   in
		want int
	}{
		{in: in{m1: 11, d1: 16, m2: 11, d2: 17}, want: 0},
		{in: in{m1: 11, d1: 30, m2: 12, d2: 1}, want: 1},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.m1, tt.in.d1, tt.in.m2, tt.in.d2)
			if got != tt.want {
				t.Fatalf("in: %v got: %v, want: %v", tt.in, got, tt.want)
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

	// Output: 1
}
