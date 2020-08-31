package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Parallel()
	type in struct {
		n, m   int
		matrix [][]int
	}
	cases := []struct {
		in   in
		want int
	}{
		{in: in{n: 3, m: 3, matrix: [][]int{{1, 7}, {3, 2}, {1, 7}}}, want: 702},
		{in: in{n: 3, m: 2, matrix: [][]int{{2, 1}, {2, 3}}}, want: -1},
		{in: in{n: 2, m: 3, matrix: [][]int{{2, 0}, {2, 0}, {2, 0}}}, want: 10},
		{in: in{n: 3, m: 0}, want: 100},
		{in: in{n: 1, m: 0}, want: 0},
	}
	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(c.in.n, c.in.m, c.in.matrix)
			if got != c.want {
				t.Errorf("in: %+v, got: %d, want: %d", c.in, got, c.want)
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

	// Output: -1
}
