package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

type Case struct {
	input, want string
}

var cases = []Case{
	{input: "sardine", want: "xxxxxxx"},
	{input: "xxxx", want: "xxxx"},
}

func TestSolution(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.input)
			if got != tt.want {
				t.Errorf("got: %s, want: %s", got, tt.want)
			}
		})
	}
}

func Example_main() {
	s := strings.NewReader("gone")
	b, _ := ioutil.ReadAll(s)
	inr, inw, _ := os.Pipe()
	orgStdin := os.Stdin
	inw.Write(b)
	inw.Close()
	os.Stdin = orgStdin
	os.Stdin = inr

	main()

	os.Stdin = orgStdin

	// Output: xxxx
}
