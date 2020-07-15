package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

type input struct {
	a, b int
}

type Case struct {
	input input
	want  int
}

var cases = []Case{
	{input: input{a: 4, b: 3}, want: 3333},
	{input: input{a: 7, b: 7}, want: 7777777},
}

func TestSolution(t *testing.T) {
	t.Parallel()
	for i, tt := range cases {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.input.a, tt.input.b)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}

func Example_main() {
	s := strings.NewReader("4 3")
	b, _ := ioutil.ReadAll(s)
	inr, inw, _ := os.Pipe()
	orgStdin := os.Stdin
	inw.Write(b)
	inw.Close()
	os.Stdin = orgStdin
	os.Stdin = inr

	main()

	os.Stdin = orgStdin

	// Output: 3333
}
