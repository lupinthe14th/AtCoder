package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestSolution(t *testing.T) {
	type in struct {
		n int
		h []int
	}
	tests := []struct {
		in   in
		want string
	}{
		{in: in{n: 5, h: []int{1, 2, 1, 1, 3}}, want: "Yes"},
		{in: in{n: 4, h: []int{1, 3, 2, 1}}, want: "No"},
		{in: in{n: 5, h: []int{1, 2, 3, 4, 5}}, want: "Yes"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in.n, tt.in.h)
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
