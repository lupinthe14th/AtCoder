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
		n  int
		ab [][2]int
	}
	tests := []struct {
		in   in
		want int64
	}{
		{in: in{n: 3, ab: [][2]int{{5, 7}, {2, 6}, {8, 10}}}, want: 18},
		{in: in{n: 5, ab: [][2]int{{1, 71}, {43, 64}, {13, 35}, {14, 54}, {79, 85}}}, want: 334},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.n, tt.in.ab)
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

	// Output: 8494550716
}
