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
		a, b, c int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{a: 3, b: 3, c: 3}, want: 9},
		{in: in{a: 2, b: 2, c: 4}, want: 0},
		{in: in{a: 2, b: 2, c: 5}, want: 0},
		{in: in{a: 5, b: 3, c: 5}, want: 15},
		{in: in{a: 5, b: 3, c: 4}, want: 0},
		{in: in{a: 5, b: 3, c: 7}, want: 15},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.a, tt.in.b, tt.in.c)
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

	// Output: 15
}
