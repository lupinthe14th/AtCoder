package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestSolution(t *testing.T) {
	type in struct {
		a, b int
	}

	var cases = []struct {
		in   in
		want string
	}{
		{in: in{a: 1, b: 21}, want: "Yes"},
		{in: in{a: 12, b: 10}, want: "No"},
	}
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in.a, tt.in.b)
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

	// Output: No
}
