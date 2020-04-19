package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type in struct {
		n    int
		nums []int
	}
	var cases = []struct {
		in   in
		want []int
	}{
		{in: in{n: 5, nums: []int{1, 1, 2, 2}}, want: []int{2, 2, 0, 0, 0}},
		{in: in{n: 7, nums: []int{1, 2, 3, 4, 5, 6}}, want: []int{1, 1, 1, 1, 1, 1, 0}},
	}
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in.n, tt.in.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("in: %+v, got: %v, want: %v", tt.in, got, tt.want)
			}
		})
	}
}

func Example_main() {
	c, _ := ioutil.ReadFile("./input.txt")
	inr, inw, _ := os.Pipe()
	orgStdin := os.Stdin
	inw.Write(c)
	os.Stdin = orgStdin
	defer func() {
		inw.Close()
		os.Stdin = orgStdin
	}()
	os.Stdin = inr

	main()

	// Output:
	// 9
	// 0
	// 0
	// 0
	// 0
	// 0
	// 0
	// 0
	// 0
	// 0
}
