package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type input struct {
	n int
	s string
}

type Case struct {
	input input
	want  int
}

var cases = []Case{
	{input: input{n: 10, s: "ZABCDBABCQ"}, want: 2},
}

func TestSolution(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.input.n, tt.input.s)
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
	defer func() { os.Stdin = orgStdin }()
	os.Stdin = inr

	main()

	// Output: 0
}
