package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type Case struct {
	input input
	want  float64
}

type input struct {
	n, k int
	p    []int
}

var cases = []Case{
	{input: input{n: 5, k: 3, p: []int{1, 2, 2, 4, 5}}, want: 7.000000000000},
	{input: input{n: 10, k: 4, p: []int{17, 13, 13, 12, 15, 20, 10, 13, 17, 11}}, want: 32.000000000000},
	{input: input{n: 1, k: 1, p: []int{6}}, want: 3.5},
}

func TestSolutio(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.input.n, tt.input.k, tt.input.p)
			if got != tt.want {
				t.Errorf("%f, want: %f", got, tt.want)
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

	// Output: 3.500000000000
}
