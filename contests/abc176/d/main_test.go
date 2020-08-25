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
		c, d [2]int
		s    []string
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{h: 4, w: 4, c: [2]int{1, 1}, d: [2]int{4, 4}, s: []string{"..#.", "..#.", ".#..", ".#.."}}, want: 1},
		{in: in{h: 4, w: 4, c: [2]int{1, 4}, d: [2]int{4, 1}, s: []string{".##.", "####", "####", ".##."}}, want: -1},
		{in: in{h: 4, w: 4, c: [2]int{2, 2}, d: [2]int{3, 3}, s: []string{"....", "....", "....", "...."}}, want: 0},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.h, tt.in.w, tt.in.c, tt.in.d, tt.in.s)
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

	// Output: 2
}
