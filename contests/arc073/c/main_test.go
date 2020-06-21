package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestSolution(t *testing.T) {
	type in struct {
		n, t int
		ts   []int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{n: 2, t: 4, ts: []int{0, 3}}, want: 7},
		{in: in{n: 2, t: 4, ts: []int{0, 5}}, want: 8},
		{in: in{n: 4, t: 1000000000, ts: []int{0, 1000, 1000000, 1000000000}}, want: 2000000000},
		{in: in{n: 1, t: 1, ts: []int{0}}, want: 1},
		{in: in{n: 9, t: 10, ts: []int{0, 3, 5, 7, 100, 110, 200, 300, 311}}, want: 67},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in.n, tt.in.t, tt.in.ts)
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

	// Output: 2000000000
}
