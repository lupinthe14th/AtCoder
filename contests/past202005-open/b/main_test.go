package main

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Parallel()
	type in struct {
		n, m, q int
		s       [][]int
	}
	tests := []struct {
		in   in
		want []int
	}{
		{in: in{n: 2, m: 1, q: 6, s: [][]int{{2, 1, 1}, {1, 1}, {1, 2}, {2, 2, 1}, {1, 1}, {1, 2}}}, want: []int{1, 0, 0, 0}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.n, tt.in.m, tt.in.q, tt.in.s)
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
	// 0
	// 4
	// 3
	// 0
	// 3
	// 10
	// 9
	// 4
	// 4
	// 4
	// 0
	// 0
	// 9
	// 3
	// 9
	// 0
	// 3
}
