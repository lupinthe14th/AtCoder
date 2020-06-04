package main

import (
	"fmt"
	"os"
	"testing"
)

func TestSolution(t *testing.T) {
	type in struct {
		a, b int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{a: 5, b: 3}, want: 9},
		{in: in{a: 3, b: 4}, want: 7},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in.a, tt.in.b)
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

	// Output: 12
}