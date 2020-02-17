package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type in struct {
	n, k int
	nums []int
}

type Case struct {
	in   in
	want int64
}

var cases = []Case{
	{in: in{n: 4, k: 3, nums: []int{3, 3, -4, -2}}, want: -6},
	{in: in{n: 10, k: 40, nums: []int{5, 4, 3, 2, -1, 0, 0, 0, 0, 0}}, want: 6},
	{in: in{n: 10, k: 35, nums: []int{5, 4, 3, 2, -1, 0, 0, 0, 0, 0}}, want: 0},
}

func TestSolution(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in.n, tt.in.k, tt.in.nums)
			if got != tt.want {
				t.Errorf("in: %+v, got: %d, want: %d", tt.in, got, tt.want)
			}
		})
	}
}

func Example_main() {
	c, _ := ioutil.ReadFile("./input.txt")
	inr, inw, _ := os.Pipe()
	orgStdin := os.Stdin
	inw.Write(c)
	inw.Close()
	os.Stdin = orgStdin
	defer func() { os.Stdin = orgStdin }()
	os.Stdin = inr

	main()

	// Output: 448283280358331064
}
