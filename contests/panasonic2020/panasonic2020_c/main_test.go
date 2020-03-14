package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type in struct {
	a, b, c int64
}

type Case struct {
	in   in
	want string
}

var cases = []Case{
	{in: in{a: 2, b: 3, c: 9}, want: "No"},
	{in: in{a: 249999999, b: 250000000, c: 999999998}, want: "Yes"},
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

	// Output: Yes
}
