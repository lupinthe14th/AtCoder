package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type Case struct {
	in   int
	want int
}

var cases = []Case{
	{in: 5, want: 3},
	{in: 2, want: 1},
	{in: 100, want: 50},
}

func TestSolution(t *testing.T) {
	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(c.in)
			if got != c.want {
				t.Errorf("in: %+v, got: %d, want: %d", c.in, got, c.want)
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

	// Output: 1
}
