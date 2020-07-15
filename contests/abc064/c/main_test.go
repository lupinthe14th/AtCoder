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
		{in: in{n: 4, a: []int{2100, 2500, 2700, 2700}}, want: []int{2, 2}},
		{in: in{n: 5, a: []int{1100, 1900, 2800, 3200, 3200}}, want: []int{3, 5}},
		{in: in{n: 5, a: []int{3200, 3200, 3200, 3200, 3200}}, want: []int{1, 5}},
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

	// Output: 1 1
}
