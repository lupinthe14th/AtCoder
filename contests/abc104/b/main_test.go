package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{in: "AtCoder", want: "AC"},
		{in: "AcycliC", want: "WA"},
		{in: "AtCoCo", want: "WA"},
		{in: "Atcoder", want: "WA"},
		{in: "atcoder", want: "WA"},
		{in: "AtCOder", want: "WA"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
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

	// Output: WA
}
