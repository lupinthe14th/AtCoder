package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Parallel()
	type in struct {
		a, b, c, x, y int
	}

	var cases = []struct {
		in   in
		want int
	}{
		{in: in{a: 1500, b: 2000, c: 1600, x: 3, y: 2}, want: 7900},
		{in: in{a: 1500, b: 2000, c: 1900, x: 3, y: 2}, want: 8500},
	}
	for i, tt := range cases {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.a, tt.in.b, tt.in.c, tt.in.x, tt.in.y)
			if got != tt.want {
				t.Errorf("in: %+v, got: %v, want: %v", tt.in, got, tt.want)
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

	// Output: 100000000
}
