package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestSolution(t *testing.T) {
	type in struct {
		n    int
		a, b []int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{n: 2, a: []int{3, 5, 2}, b: []int{4, 5}}, want: 9},
		{in: in{n: 2, a: []int{100, 1, 1}, b: []int{1, 100}}, want: 3},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in.n, tt.in.a, tt.in.b)
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

	// Output: 22
}
