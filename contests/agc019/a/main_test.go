package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestSolution(t *testing.T) {
	type in struct {
		q, h, s, d, n int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{q: 20, h: 30, s: 70, d: 90, n: 3}, want: 150},
		{in: in{q: 10000, h: 1000, s: 100, d: 10, n: 1}, want: 100},
		{in: in{q: 10, h: 100, s: 1000, d: 10000, n: 1}, want: 40},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in.q, tt.in.h, tt.in.s, tt.in.d, tt.in.n)
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

	// Output: 1524157763907942
}
