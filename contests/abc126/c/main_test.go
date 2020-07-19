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
		n, k int
	}
	tests := []struct {
		in   in
		want float64
	}{
		{in: in{n: 3, k: 10}, want: 0.14583333333333331},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.n, tt.in.k)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
func TestHelper(t *testing.T) {
	t.Parallel()
	type in struct {
		n, k int
	}
	tests := []struct {
		in   in
		want float64
	}{
		{in: in{n: 1, k: 10}, want: 4.0},
		{in: in{n: 2, k: 10}, want: 3.0},
		{in: in{n: 3, k: 10}, want: 2.0},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := helper(tt.in.n, tt.in.k)
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

	// Output: 0.999973749998
}
