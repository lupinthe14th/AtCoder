package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestSolution(t *testing.T) {
	type in struct {
		n    int
		nums []int
	}
	var cases = []struct {
		in   in
		want int
	}{
		{in: in{n: 41, nums: []int{5, 6}}, want: 30},
		{in: in{n: 10, nums: []int{5, 6}}, want: -1},
		{in: in{n: 11, nums: []int{5, 6}}, want: 0},
		{in: in{n: 314, nums: []int{9, 26, 5, 35, 8, 9, 79, 3, 23, 8, 46, 2, 6, 43, 3}}, want: 9},
	}
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in.n, tt.in.nums)
			if got != tt.want {
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

	// Output: -1
}
