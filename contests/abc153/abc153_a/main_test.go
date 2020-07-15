package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

type input struct {
	h, a int
}

type Case struct {
	input input
	want  int
}

var cases = []Case{
	{input: input{h: 10, a: 4}, want: 3},
	{input: input{h: 1, a: 10000}, want: 1},
	{input: input{h: 10000, a: 1}, want: 10000},
}

func TestSolution(t *testing.T) {
	t.Parallel()
	for i, tt := range cases {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.input.h, tt.input.a)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}

func Example_main() {
	s := strings.NewReader("10 4")
	b, _ := ioutil.ReadAll(s)
	inr, inw, _ := os.Pipe()
	orgStdin := os.Stdin
	inw.Write(b)
	inw.Close()
	os.Stdin = orgStdin
	os.Stdin = inr

	main()

	os.Stdin = orgStdin

	// Output: 3
}
