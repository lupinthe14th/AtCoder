package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		in, want int
		wanterr  string
	}{
		{in: 432, want: 400},
		{in: 1079, wanterr: ":("},
		{in: 1001, want: 927},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got, err := solution(tt.in)
			if got == 0 && err != tt.wanterr {
				t.Fatalf("in: %v got: %v, want: %v", tt.in, err, tt.wanterr)
			}
			if got != tt.want {
				t.Fatalf("in: %v, got: %v want: %v", tt.in, got, tt.want)
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

	// Output: 927
}
