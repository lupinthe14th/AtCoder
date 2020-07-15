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
		h, w, k int
		c       [][]byte
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{h: 2, w: 3, k: 2, c: [][]byte{{'.', '.', '#'}, {'#', '#', '#'}}}, want: 5},
		{in: in{h: 2, w: 3, k: 4, c: [][]byte{{'.', '.', '#'}, {'#', '#', '#'}}}, want: 1},
		{in: in{h: 2, w: 2, k: 3, c: [][]byte{{'#', '#'}, {'#', '#'}}}, want: 0},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.h, tt.in.w, tt.in.k, tt.in.c)
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

	// Output: 208
}
