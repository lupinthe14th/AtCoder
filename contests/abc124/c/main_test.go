package main

import (
	"fmt"
	"os"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{in: "000", want: 1},
		{in: "0", want: 0},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in)
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

	// Output: 3
}
