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
		n, m   int
		ab, cd [][]int
	}
	tests := []struct {
		in   in
		want []int
	}{
		{in: in{n: 2, m: 2, ab: [][]int{{2, 0}, {0, 0}}, cd: [][]int{{-1, 0}, {1, 0}}}, want: []int{2, 1}},
		{in: in{n: 3, m: 4, ab: [][]int{{10, 10}, {-10, -10}, {3, 3}}, cd: [][]int{{1, 2}, {2, 3}, {3, 5}, {3, 5}}}, want: []int{3, 1, 2}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.n, tt.in.m, tt.in.ab, tt.in.cd)
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
	// 5
	// 4
	// 3
	// 2
	// 1
}
