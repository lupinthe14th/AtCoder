package main

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type in struct {
		n int
		s []string
	}
	tests := []struct {
		in   in
		want []string
	}{
		{in: in{n: 6, s: []string{"AC", "TLE", "AC", "AC", "WA", "TLE"}}, want: []string{"AC x 3", "WA x 1", "TLE x 2", "RE x 0"}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.n, tt.in.s)
			if !reflect.DeepEqual(got, tt.want) {
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

	// Output:
	// AC x 10
	// WA x 0
	// TLE x 0
	// RE x 0
}
