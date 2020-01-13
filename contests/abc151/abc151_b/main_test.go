package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type input struct {
	n, k, m int
	a       []int
}

type Case struct {
	input input
	want  int
}

var cases = []Case{
	{input: input{n: 5, k: 10, m: 7, a: []int{8, 10, 3, 6}}, want: 8},
	{input: input{n: 4, k: 100, m: 60, a: []int{100, 100, 100}}, want: 0},
	{input: input{n: 4, k: 100, m: 60, a: []int{0, 0, 0}}, want: -1},
}

func TestSolution(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.input.n, tt.input.k, tt.input.m, tt.input.a)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}

func Example_main() {
	b, _ := ioutil.ReadFile("input.txt")
	inr, inw, _ := os.Pipe()
	orgStdin := os.Stdin
	inw.Write(b)
	inw.Close()
	os.Stdin = inr

	main()

	os.Stdin = orgStdin

	// Output: 8
}
