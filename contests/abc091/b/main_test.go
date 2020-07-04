package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestSolution(t *testing.T) {
	type in struct {
		n, m int
		s, t []string
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{n: 3, s: []string{"apple", "orange", "apple"}, m: 1, t: []string{"grape"}}, want: 2},
		{in: in{n: 3, s: []string{"apple", "orange", "apple"}, m: 5, t: []string{"apple", "apple", "apple", "apple", "apple"}}, want: 1},
		{in: in{n: 1, s: []string{"voldemort"}, m: 10, t: []string{"voldemort", "voldemort", "voldemort", "voldemort", "voldemort", "voldemort", "voldemort", "voldemort", "voldemort", "voldemort"}}, want: 0},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.n, tt.in.m, tt.in.s, tt.in.t)
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

	// Output: 1
}
