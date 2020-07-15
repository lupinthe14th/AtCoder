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
		p [][2]int
	}
	tests := []struct {
		in   in
		want float64
	}{
		{in: in{n: 3, p: [][2]int{{0, 0}, {1, 0}, {0, 1}}}, want: 2.2761423749153966},
		{in: in{n: 2, p: [][2]int{{-879, 981}, {-866, 890}}}, want: 91.92388155425118},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.n, tt.in.p)
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

	// Output: 7641.9817824387
}
