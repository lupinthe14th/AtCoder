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
		n    int
		a, b []int
	}
	tests := []struct {
		in   in
		want int64
	}{
		{in: in{n: 3, a: []int{5, 2, 8}, b: []int{7, 6, 10}}, want: 18},
		{in: in{n: 5, a: []int{1, 43, 13, 14, 79}, b: []int{71, 64, 35, 54, 85}}, want: 334},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
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

	// Output: 8494550716
}
