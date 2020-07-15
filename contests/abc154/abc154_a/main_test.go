package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type input struct {
	s, t string
	a, b int
	u    string
}

type want struct {
	m, n int
}

type Case struct {
	input input
	want  want
}

var cases = []Case{
	{input: input{s: "red", t: "blue", a: 3, b: 4, u: "red"}, want: want{m: 2, n: 4}},
}

func TestSolution(t *testing.T) {
	t.Parallel()
	for i, tt := range cases {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			gotM, gotN := solution(tt.input.s, tt.input.t, tt.input.u, tt.input.a, tt.input.b)
			if gotM != tt.want.m && gotN != tt.want.n {
				t.Errorf("gotM: %d, gotN: %d, want: %v", gotM, gotN, tt.want)
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

	// Output: 5 4
}
