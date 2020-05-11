package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestSolution(t *testing.T) {
	type in struct {
		n, m, x int
		a       [][]int
		c       []int
	}

	var cases = []struct {
		in   in
		want int
	}{
		{in: in{n: 3, m: 3, x: 10, a: [][]int{{2, 2, 4}, {8, 7, 9}, {2, 3, 9}}, c: []int{60, 70, 50}}, want: 120},
		{in: in{n: 3, m: 3, x: 10, a: [][]int{{3, 1, 4}, {1, 5, 9}, {2, 6, 5}}, c: []int{100, 100, 100}}, want: -1},
	}
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in.n, tt.in.m, tt.in.x, tt.in.a, tt.in.c)
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
	inw.Close()
	os.Stdin = orgStdin
	defer func() { os.Stdin = orgStdin }()
	os.Stdin = inr

	main()

	// Output: 1067
}
