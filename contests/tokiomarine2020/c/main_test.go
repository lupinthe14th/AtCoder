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
		n, k int
		a    []int
	}
	tests := []struct {
		in   in
		want []int
	}{
		{in: in{n: 5, k: 1, a: []int{1, 0, 0, 1, 0}}, want: []int{1, 2, 2, 1, 2}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in.n, tt.in.k, tt.in.a)
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

	// Output: 3 3 4 4 3
}
