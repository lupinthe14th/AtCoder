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
		d [][2]int
	}
	tests := []struct {
		in   in
		want string
	}{
		{in: in{n: 5, d: [][2]int{{1, 2}, {6, 6}, {4, 4}, {3, 3}, {3, 2}}}, want: "Yes"},
		{in: in{n: 5, d: [][2]int{{1, 1}, {2, 2}, {3, 4}, {5, 5}, {6, 6}}}, want: "No"},
		{in: in{n: 6, d: [][2]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}}}, want: "Yes"},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.n, tt.in.d)
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

	// Output: Yes
}
