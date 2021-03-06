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
		s, t string
	}
	tests := []struct {
		in   in
		want string
	}{
		{in: in{s: "yx", t: "axy"}, want: "Yes"},
		{in: in{s: "ratcode", t: "atlas"}, want: "Yes"},
		{in: in{s: "cd", t: "abc"}, want: "No"},
		{in: in{s: "w", t: "ww"}, want: "Yes"},
		{in: in{s: "zzz", t: "zzz"}, want: "No"},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.s, tt.in.t)
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
