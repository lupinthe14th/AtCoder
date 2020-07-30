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
		s string
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{n: 4, s: "0224"}, want: 3},
		{in: in{n: 6, s: "123123"}, want: 17},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			// t.Parallel()
			got := solution(tt.in.n, tt.in.s)
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

	// Output: 329
}
