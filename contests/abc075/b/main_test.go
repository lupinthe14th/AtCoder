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
		h, w int
		s    []string
	}
	tests := []struct {
		in   in
		want []string
	}{
		{in: in{h: 3, w: 5, s: []string{".....", ".#.#.", "....."}}, want: []string{"11211", "1#2#1", "11211"}},
		{in: in{h: 3, w: 5, s: []string{"#####", "#####", "#####"}}, want: []string{"#####", "#####", "#####"}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in.h, tt.in.w, tt.in.s)
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
	// #####3
	// #8#7##
	// ####5#
	// 4#65#2
	// #5##21
	// #4#310

}
