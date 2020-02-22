package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type in struct {
	n, r int
}

type Case struct {
	in   in
	want int
}

var cases = []Case{
	{in: in{n: 2, r: 2919}, want: 3719},
}

func TestSolution(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in.n, tt.in.r)
			if got != tt.want {
				t.Errorf("in: %+v, got: %d, want: %d", tt.in, got, tt.want)
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

	// Output: 3051
}
