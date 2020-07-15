package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type input struct {
	h, n int
	a    []int
}

type Case struct {
	input input
	want  string
}

var cases = []Case{
	{input: input{h: 10, n: 3, a: []int{4, 5, 6}}, want: "Yes"},
	{input: input{h: 20, n: 3, a: []int{4, 5, 6}}, want: "No"},
	{input: input{h: 210, n: 5, a: []int{31, 41, 59, 26, 53}}, want: "Yes"},
	{input: input{h: 211, n: 5, a: []int{31, 41, 59, 26, 53}}, want: "No"},
}

func TestSolution(t *testing.T) {
	t.Parallel()
	for i, tt := range cases {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.input.h, tt.input.n, tt.input.a)
			if got != tt.want {
				t.Errorf("%s, want: %s", got, tt.want)
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

	// Output: Yes
}
