package main

import (
	"fmt"
	"os"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		in   []int
		want string
	}{
		{in: []int{1, 2, 4, 8, 9, 15}, want: "Yay!"},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in)
			if got != tt.want {
				t.Fatalf("in: %+v, got: %v, want: %v", tt.in, got, tt.want)
			}
		})
	}
}

func Example_main() {
	fd, _ := os.Open("./input.txt")
	orgStdin := os.Stdin
	os.Stdin = orgStdin
	os.Stdin = fd
	defer func() {
		os.Stdin = orgStdin
		fd.Close()
	}()

	main()

	// Output: :(
}
