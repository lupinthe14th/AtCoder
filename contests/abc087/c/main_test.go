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
		n int
		a [2][]int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{n: 5, a: [2][]int{{3, 2, 2, 4, 1}, {1, 2, 2, 2, 1}}}, want: 14},
		{in: in{n: 4, a: [2][]int{{1, 1, 1, 1}, {1, 1, 1, 1}}}, want: 5},
		{in: in{n: 1, a: [2][]int{{2}, {3}}}, want: 5},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.n, tt.in.a)
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

	// Output: 29
}
