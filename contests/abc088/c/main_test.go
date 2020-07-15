package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in   [3][3]int
		want string
	}{
		{in: [3][3]int{{1, 0, 1}, {2, 1, 2}, {1, 0, 1}}, want: "Yes"},
		{in: [3][3]int{{2, 2, 2}, {2, 1, 2}, {2, 2, 2}}, want: "No"},
		{in: [3][3]int{{0, 8, 8}, {0, 8, 8}, {0, 8, 8}}, want: "Yes"},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in)
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

	// Output: No
}
