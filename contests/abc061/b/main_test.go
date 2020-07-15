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
		n, m int
		ab   [][]int
	}
	tests := []struct {
		in   in
		want []int
	}{
		{in: in{n: 4, m: 3, ab: [][]int{{1, 2}, {2, 3}, {1, 4}}}, want: []int{2, 2, 1, 1}},
		{in: in{n: 2, m: 5, ab: [][]int{{1, 2}, {2, 1}, {1, 2}, {2, 1}, {1, 2}}}, want: []int{5, 5}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.n, tt.in.m, tt.in.ab)
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
	// 3
	// 3
	// 2
	// 2
	// 2
	// 1
	// 1
	// 2
}
