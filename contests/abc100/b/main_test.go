package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestSolution(t *testing.T) {
	type in struct {
		d, n int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{d: 0, n: 5}, want: 5},
		{in: in{d: 0, n: 100}, want: 101},
		{in: in{d: 1, n: 11}, want: 1100},
	}
	for i, tt := range tests {
		i := i   // capture range variable
		tt := tt // capture range variable
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.d, tt.in.n)
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

	// Output: 850000
}
