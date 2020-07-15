package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type in struct {
	a, b []int
	xs   [][]int
}

type Case struct {
	in   in
	want int
}

var cases = []Case{
	{in: in{a: []int{3, 3}, b: []int{3, 3, 3}, xs: [][]int{{1, 2, 1}}}, want: 5},
	{in: in{a: []int{10}, b: []int{10}, xs: [][]int{{1, 1, 5}, {1, 1, 10}}}, want: 10},
}

func TestSolution(t *testing.T) {
	t.Parallel()
	for i, tt := range cases {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.a, tt.in.b, tt.in.xs)
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

	// Output: 6
}
