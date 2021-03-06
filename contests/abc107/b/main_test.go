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
		h, w int
		a    []string
	}
	tests := []struct {
		in   in
		want [][]byte
	}{
		{in: in{h: 4, w: 4, a: []string{"##.#", "....", "##.#", ".#.#"}}, want: [][]byte{{'#', '#', '#'}, {'#', '#', '#'}, {'.', '#', '#'}}},
		{in: in{h: 3, w: 3, a: []string{"#..", ".#.", "..#"}}, want: [][]byte{{'#', '.', '.'}, {'.', '#', '.'}, {'.', '.', '#'}}},
		{in: in{h: 4, w: 5, a: []string{".....", ".....", "..#..", "....."}}, want: [][]byte{{'#'}}},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.h, tt.in.w, tt.in.a)
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
	// ..#
	// #..
	// .#.
	// .#.
	// #.#
}
