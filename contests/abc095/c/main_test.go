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
		a, b, c, x, y int
	}

	var tests = []struct {
		in   in
		want int
	}{
		{in: in{a: 1500, b: 2000, c: 1600, x: 3, y: 2}, want: 7900},
		{in: in{a: 1500, b: 2000, c: 1900, x: 3, y: 2}, want: 8500},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.a, tt.in.b, tt.in.c, tt.in.x, tt.in.y)
			if got != tt.want {
				t.Fatalf("in: %+v, got: %v, want: %v", tt.in, got, tt.want)
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

	// Output: 100000000
}
