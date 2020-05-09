package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestSolution(t *testing.T) {
	type in struct {
		n, m, c int
		a       [][]int
		b       []int
	}

	var cases = []struct {
		in   in
		want int
	}{
		{in: in{n: 2, m: 3, c: -10, a: [][]int{{3, 2, 1}, {1, 2, 2}}, b: []int{1, 2, 3}}, want: 1},
		{in: in{n: 3, m: 3, c: 0, a: [][]int{{0, 100, 100}, {100, 100, 100}, {-100, 100, 100}}, b: []int{100, -100, 0}}, want: 0},
	}
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in.n, tt.in.m, tt.in.c, tt.in.a, tt.in.b)
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

	// Output: 2
}
