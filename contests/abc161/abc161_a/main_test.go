package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestSolution(t *testing.T) {
	type in struct {
		x, y, z int
	}

	type out struct {
		a, b, c int
	}
	var cases = []struct {
		in   in
		want out
	}{
		{in: in{x: 1, y: 2, z: 3}, want: out{a: 3, b: 1, c: 2}},
		{in: in{x: 41, y: 59, z: 31}, want: out{a: 31, b: 41, c: 59}},
	}
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			a, b, c := solution(tt.in.x, tt.in.y, tt.in.z)
			if a != tt.want.a && b != tt.want.b && c != tt.want.c {
				t.Errorf("in: %+v, a: %d, b: %d, c: %d, want: %d", tt.in, a, b, c, tt.want)
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

	// Output: 100 100 100
}
