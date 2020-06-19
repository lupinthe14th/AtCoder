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
		c    []string
	}
	tests := []struct {
		in   in
		want []string
	}{
		{in: in{h: 2, w: 2, c: []string{"*.", ".*"}}, want: []string{"*.", "*.", ".*", ".*"}},
		{in: in{h: 1, w: 4, c: []string{"***."}}, want: []string{"***.", "***."}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in.h, tt.in.w, tt.in.c)
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
	// .....***....***.....
	// .....***....***.....
	// ....*...*..*...*....
	// ....*...*..*...*....
	// ...*.....**.....*...
	// ...*.....**.....*...
	// ...*.....*......*...
	// ...*.....*......*...
	// ....*.....*....*....
	// ....*.....*....*....
	// .....**..*...**.....
	// .....**..*...**.....
	// .......*..*.*.......
	// .......*..*.*.......
	// ........**.*........
	// ........**.*........
	// .........**.........
	// .........**.........
}
