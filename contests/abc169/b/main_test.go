package main

import (
	"fmt"
	"os"
	"testing"
)

func TestSolution(t *testing.T) {
	type in struct {
		n int
		a []int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{n: 2, a: []int{1000000000, 1000000000}}, want: 1000000000000000000},
		{in: in{n: 3, a: []int{101, 9901, 999999000001}}, want: -1},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in.n, tt.in.a)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}

func Example_main() {
	fd, _ := os.Open("./input.txt")
	orgStdin := os.Stdin
	os.Stdin = fd
	defer func() {
		os.Stdin = orgStdin
		fd.Close()
	}()

	main()

	// Output: 0
}