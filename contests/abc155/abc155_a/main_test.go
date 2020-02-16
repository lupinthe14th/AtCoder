package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type in struct {
	a, b, c int
}

type Case struct {
	in   in
	want string
}

var cases = []Case{
	{in: in{a: 5, b: 7, c: 5}, want: "Yes"},
	{in: in{a: 5, b: 7, c: 7}, want: "Yes"},
	{in: in{a: 5, b: 5, c: 7}, want: "Yes"},
}

func TestSolution(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in.a, tt.in.b, tt.in.c)
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

	// Output: No
}
