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
	{in: "akasaka", want: "Yes"},
	{in: "atcoder", want: "No"},
}

func TestSolution(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in)
			if got != tt.want {
				t.Errorf("in: %+v, got: %v, want: %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	var cases = []struct {
		in   string
		want bool
	}{
		{in: "akasaka", want: true},
		{in: "atcoder", want: false},
	}
	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := isPalindrome(c.in)
			if got != c.want {
				t.Errorf("got: %v, want: %v", got, c.want)
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