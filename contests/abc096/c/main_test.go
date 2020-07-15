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
		h, w int
		s    [][]string
	}
	tests := []struct {
		in   in
		want string
	}{
		{in: in{h: 3, w: 3, s: [][]string{{".", "#", "."}, {"#", "#", "#"}, {".", "#", "."}}}, want: "Yes"},
		{in: in{h: 5, w: 5, s: [][]string{
			{"#", ".", "#", ".", "#"},
			{".", "#", ".", "#", "."},
			{"#", ".", "#", ".", "#"},
			{".", "#", ".", "#", "."},
			{"#", ".", "#", ".", "#"},
		}}, want: "No"},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.h, tt.in.w, tt.in.s)
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
