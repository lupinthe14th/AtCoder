package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestSolution(t *testing.T) {
	type in struct {
		n, m int
		lr   [][]int
	}

	var cases = []struct {
		in   in
		want int
	}{
		{in: in{n: 4, m: 2, lr: [][]int{{1, 3}, {2, 4}}}, want: 2},
		{in: in{n: 10, m: 3, lr: [][]int{{3, 6}, {5, 7}, {6, 9}}}, want: 1},
		{in: in{n: 100000, m: 1, lr: [][]int{{1, 100000}}}, want: 100000},
	}
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in.n, tt.in.m, tt.in.lr)
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

	// Output: 0
}
