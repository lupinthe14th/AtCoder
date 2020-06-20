package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestSolution(t *testing.T) {
	type in struct {
		n int
		s []int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{n: 3, s: []int{5, 10, 15}}, want: 25},
		{in: in{n: 3, s: []int{10, 10, 15}}, want: 35},
		{in: in{n: 4, s: []int{10, 10, 15, 25}}, want: 45},
		{in: in{n: 1, s: []int{10}}, want: 0},
		{in: in{n: 1, s: []int{1}}, want: 1},
		{in: in{n: 1, s: []int{100}}, want: 0},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in.n, tt.in.s)
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

	// Output: 0
}
