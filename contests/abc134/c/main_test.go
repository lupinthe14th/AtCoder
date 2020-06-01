package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type in struct {
		n int
		a []int
	}

	var cases = []struct {
		in   in
		want []int
	}{
		{in: in{n: 3, a: []int{1, 4, 3}}, want: []int{4, 3, 4}},
		{in: in{n: 2, a: []int{200000, 1}}, want: []int{1, 200000}},
	}
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in.n, tt.in.a)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("in: %+v, got: %v, want: %v", tt.in, got, tt.want)
			}
		})
	}
}

func Example_main() {
	fd, _ := os.Open("./input.txt")
	orgStdin := os.Stdin
	os.Stdin = orgStdin
	os.Stdin = fd
	defer func() {
		os.Stdin = orgStdin
		fd.Close()
	}()

	main()

	// Output:
	// 5
	// 5
}
