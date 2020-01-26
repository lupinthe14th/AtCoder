package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type input struct {
	n, k int
	h    []int
}

type Case struct {
	input input
	want  int
}

var cases = []Case{
	{input: input{n: 3, k: 1, h: []int{4, 1, 5}}, want: 5},
	{input: input{n: 8, k: 9, h: []int{7, 9, 3, 2, 3, 8, 4, 6}}, want: 0},
	{input: input{n: 3, k: 0, h: []int{1000000000, 1000000000, 1000000000}}, want: 3000000000},
}

func TestSolution(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.input.n, tt.input.k, tt.input.h)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
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
	os.Stdin = inr

	main()

	os.Stdin = orgStdin

	// Output: 5
}
