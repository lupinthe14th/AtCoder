package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type input struct {
	k, x int
}

type Case struct {
	input input
	want  bool
}

var cases = []Case{
	{input: input{k: 2, x: 900}, want: true},
	{input: input{k: 1, x: 501}, want: false},
}

func TestSolution(t *testing.T) {
	t.Parallel()
	for i, tt := range cases {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.input.k, tt.input.x)
			if got != tt.want {
				t.Errorf("%t, want: %t", got, tt.want)
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

	// Output: Yes
}

func Example_main2() {
	c, _ := ioutil.ReadFile("./input2.txt")
	inr, inw, _ := os.Pipe()
	orgStdin := os.Stdin
	inw.Write(c)
	inw.Close()
	os.Stdin = orgStdin
	defer func() { os.Stdin = orgStdin }()
	os.Stdin = inr

	main()

	// Output: No
}
