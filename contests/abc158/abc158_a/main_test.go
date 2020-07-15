package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type Case struct {
	in   string
	want string
}

var cases = []Case{
	{in: "ABA", want: "Yes"},
	{in: "BBA", want: "Yes"},
	{in: "BBB", want: "No"},
}

func TestSolution(t *testing.T) {
	t.Parallel()
	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(c.in)
			if got != c.want {
				t.Errorf("in: %+v, got: %s, want: %s", c.in, got, c.want)
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
