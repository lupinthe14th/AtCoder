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
		n int
		a []int
	}
	tests := []struct {
		in   in
		want []int
	}{
		{in: in{n: 5, a: []int{1, 1, 2, 1, 2}}, want: []int{2, 2, 3, 2, 3}},
		{in: in{n: 4, a: []int{1, 2, 3, 4}}, want: []int{0, 0, 0, 0}},
		{in: in{n: 5, a: []int{3, 3, 3, 3, 3}}, want: []int{6, 6, 6, 6, 6}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.n, tt.in.a)
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
	// 7
	// 5
	// 7
	// 7
	// 5
	// 7
	// 5
}
