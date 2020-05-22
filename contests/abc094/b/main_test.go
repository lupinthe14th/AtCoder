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
		a       []int
	}

	var cases = []struct {
		in   in
		want int
	}{
		{in: in{n: 5, m: 3, x: 3, a: []int{1, 2, 4}}, want: 1},
		{in: in{n: 10, m: 7, x: 5, a: []int{1, 2, 3, 4, 6, 8, 9}}, want: 3},
	}
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in.n, tt.in.m, tt.in.x, tt.in.a)
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
