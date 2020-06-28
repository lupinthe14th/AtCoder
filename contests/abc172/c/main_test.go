package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestSolution(t *testing.T) {
	type in struct {
		n, m, k int
		a, b    []int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{n: 3, m: 4, k: 240, a: []int{60, 90, 120}, b: []int{80, 150, 80, 150}}, want: 3},
		{in: in{n: 3, m: 4, k: 730, a: []int{60, 90, 120}, b: []int{80, 150, 80, 150}}, want: 7},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in.n, tt.in.m, tt.in.k, tt.in.a, tt.in.b)
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
