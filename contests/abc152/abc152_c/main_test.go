package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type input struct {
	n int
	p []int
}

type Case struct {
	input input
	want  int
}

var cases = []Case{
	{input: input{n: 5, p: []int{4, 2, 5, 1, 3}}, want: 3},
	{input: input{n: 4, p: []int{4, 3, 2, 1}}, want: 4},
	{input: input{n: 6, p: []int{1, 2, 3, 4, 5, 6}}, want: 1},
}

func TestSolution(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.input.n, tt.input.p)
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

	// Output: 3
}
