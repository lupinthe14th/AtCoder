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
		n, m, k int
		ab, cd  [][2]int
	}
	tests := []struct {
		in   in
		want []int
	}{
		{in: in{n: 4, m: 4, k: 1, ab: [][2]int{{2, 1}, {1, 3}, {3, 2}, {3, 4}}, cd: [][2]int{{4, 1}}}, want: []int{0, 1, 0, 1}},
		{in: in{n: 5, m: 10, k: 0, ab: [][2]int{{1, 2}, {1, 3}, {1, 4}, {1, 5}, {3, 2}, {2, 4}, {2, 5}, {4, 3}, {5, 3}, {4, 5}}, cd: [][2]int{}}, want: []int{0, 0, 0, 0, 0}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.n, tt.in.m, tt.in.k, tt.in.ab, tt.in.cd)
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

	// Output: 1 3 5 4 3 3 3 3 1 0
}
