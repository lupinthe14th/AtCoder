package main

import (
	"fmt"
	"os"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Parallel()
	type in struct {
		a int
		b string
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{a: 198, b: "1.10"}, want: 217},
		{in: in{a: 1, b: "0.01"}, want: 0},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
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

	// Output: 9990000000000000
}
