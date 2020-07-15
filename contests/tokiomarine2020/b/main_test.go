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
		a, v, b, w, t int
	}
	tests := []struct {
		in   in
		want string
	}{
		{in: in{a: 1, v: 2, b: 3, w: 1, t: 3}, want: "YES"},
		{in: in{a: 1, v: 2, b: 3, w: 2, t: 3}, want: "NO"},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.a, tt.in.v, tt.in.b, tt.in.w, tt.in.t)
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
