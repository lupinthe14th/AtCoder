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
		n, m int
		a    [][]int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{n: 1, m: 2, a: [][]int{{80, 84}}}, want: 84},
		{in: in{n: 3, m: 4, a: [][]int{{37, 29, 70, 41}, {85, 69, 76, 50}, {53, 10, 95, 100}}}, want: 250},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.n, tt.in.m, tt.in.a)
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

	// Output: 581000000
}
