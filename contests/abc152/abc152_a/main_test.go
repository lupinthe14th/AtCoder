package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

type input struct {
	n, m int
}

type Case struct {
	input input
	want  string
}

var cases = []Case{
	{input: input{n: 3, m: 3}, want: "Yes"},
}

func TestSolution(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.input.n, tt.input.m)
			if got != tt.want {
				t.Errorf("%s, want: %s", got, tt.want)
			}
		})
	}
}

func Example_main() {
	s := strings.NewReader("3 3")
	b, _ := ioutil.ReadAll(s)
	inr, inw, _ := os.Pipe()
	orgStdin := os.Stdin
	inw.Write(b)
	inw.Close()
	os.Stdin = orgStdin
	os.Stdin = inr

	main()

	os.Stdin = orgStdin

	// Output: Yes
}
