package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

type input struct {
	a, b, c, d, e, f int
}
type Case struct {
	input input
	want  int
}

var cases = []Case{
	{input: input{a: 4, b: 18, c: 25, d: 20, e: 9, f: 13}, want: 18},
}

func TestSolution(t *testing.T) {
	t.Parallel()
	for i, tt := range cases {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.input.a, tt.input.b, tt.input.c, tt.input.d, tt.input.e, tt.input.f)
			if got != tt.want {
				t.Errorf("%d, want.w: %d", got, tt.want)
			}
		})
	}
}

func Example_main() {
	s := strings.NewReader("95 96 97 98 99 100")
	b, _ := ioutil.ReadAll(s)
	inr, inw, _ := os.Pipe()
	orgStdin := os.Stdin
	inw.Write(b)
	inw.Close()
	os.Stdin = orgStdin
	os.Stdin = inr

	main()

	os.Stdin = orgStdin

	// Output: 98
}
