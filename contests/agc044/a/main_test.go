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
		t      int
		matrix [][]int
	}

	var cases = []struct {
		in   in
		want []int64
	}{
		{in: in{t: 5, matrix: [][]int{{11, 1, 2, 4, 8}, {11, 1, 2, 2, 8}, {32, 10, 8, 5, 4}, {29384293847243, 454353412, 332423423, 934923490, 1}, {900000000000000000, 332423423, 454353412, 934923490, 987654321}}}, want: []int64{20, 19, 26, 3821859835, 23441258666}},
	}
	for i, tt := range cases {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.t, tt.in.matrix)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("in: %+v, got: %v, want: %v", tt.in, got, tt.want)
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
	// 20
	// 19
	// 26
	// 3821859835
	// 23441258666
}
