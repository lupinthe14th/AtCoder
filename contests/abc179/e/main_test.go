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
		n, x, m int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{n: 6, x: 2, m: 1001}, want: 1369},
		{in: in{n: 1000, x: 2, m: 16}, want: 6},
		{in: in{n: 1000, x: 2, m: 449}, want: 233423},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.n, tt.in.x, tt.in.m)
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

	// Output: 492443256176507
}
