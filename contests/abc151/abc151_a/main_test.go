package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

type Case struct {
	input string
	want  string
}

var cases = []Case{
	{input: "a", want: "b"},
	{input: "y", want: "z"},
}

func TestSolution(t *testing.T) {
	t.Parallel()
	for i, tt := range cases {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.input)
			if got != tt.want {
				t.Errorf("%s, want: %s", got, tt.want)
			}
		})
	}
}

func Example_main() {
	s := strings.NewReader("c")
	b, _ := ioutil.ReadAll(s)
	inr, inw, _ := os.Pipe()
	orgStdin := os.Stdin
	inw.Write(b)
	inw.Close()
	os.Stdin = orgStdin
	os.Stdin = inr

	main()

	os.Stdin = orgStdin

	// Output: d
}
