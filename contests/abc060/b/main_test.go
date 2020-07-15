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
		a, b, c int
	}
	tests := []struct {
		in   in
		want string
	}{
		{in: in{a: 7, b: 5, c: 1}, want: "YES"},
		{in: in{a: 2, b: 2, c: 1}, want: "NO"},
		{in: in{a: 1, b: 100, c: 97}, want: "YES"},
		{in: in{a: 40, b: 98, c: 58}, want: "YES"},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.a, tt.in.b, tt.in.c)
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

	// Output: NO
}
