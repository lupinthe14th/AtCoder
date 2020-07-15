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
	{in: "hihi", want: "Yes"},
	{in: "hihihihihi", want: "Yes"},
	{in: "hi", want: "Yes"},
	{in: "ha", want: "No"},
	{in: "hii", want: "No"},
	{in: "h", want: "No"},
}

func TestSolution(t *testing.T) {
	t.Parallel()
	for i, tt := range cases {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in)
			if got != tt.want {
				t.Errorf("in: %+v, got: %s, want: %s", tt.in, got, tt.want)
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
