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
		n, m int
		a, b []int
	}

	type want struct {
		ans  string
		nums []int
	}
	var cases = []struct {
		in   in
		want want
	}{
		{in: in{n: 4, m: 4, a: []int{1, 2, 3, 4}, b: []int{2, 3, 4, 2}}, want: want{ans: "Yes", nums: []int{1, 2, 2}}},
		{in: in{n: 6, m: 9, a: []int{3, 6, 2, 5, 4, 1, 6, 4, 5}, b: []int{4, 1, 4, 3, 6, 5, 2, 5, 6}}, want: want{ans: "Yes", nums: []int{6, 5, 6, 1, 1}}},
	}
	for i, tt := range cases {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			gotA, gotN := solution(tt.in.n, tt.in.m, tt.in.a, tt.in.b)
			if gotA != tt.want.ans || !reflect.DeepEqual(gotN, tt.want.nums) {
				t.Errorf("in: %+v, gotA: %v, gotN: %v, want: %v", tt.in, gotA, gotN, tt.want)
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
	// Yes
	// 6
	// 5
	// 6
	// 1
	// 1
}
