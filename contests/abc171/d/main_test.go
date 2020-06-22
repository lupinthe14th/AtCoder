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
		n, q    int
		a, b, c []int
	}
	tests := []struct {
		in   in
		want []uint64
	}{
		{in: in{n: 4, a: []int{1, 2, 3, 4}, q: 3, b: []int{1, 3, 2}, c: []int{2, 4, 4}}, want: []uint64{11, 12, 16}},
		{in: in{n: 4, a: []int{1, 1, 1, 1}, q: 3, b: []int{1, 2, 3}, c: []int{2, 1, 5}}, want: []uint64{8, 4, 4}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in.n, tt.in.a, tt.in.q, tt.in.b, tt.in.c)
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
	// 102
	// 200
	// 2000
}
