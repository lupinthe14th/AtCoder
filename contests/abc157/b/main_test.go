package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type in struct {
	a [][]int
	b []int
}

type Case struct {
	in   in
	want string
}

var cases = []Case{
	{in: in{a: [][]int{{84, 97, 66}, {79, 89, 11}, {61, 59, 7}}, b: []int{89, 7, 87, 79, 24, 84, 30}}, want: "Yes"},
	{in: in{a: [][]int{{31, 28, 30}, {70, 90, 84}, {86, 68, 69}}, b: []int{84, 52, 12, 69, 30, 68}}, want: "Yes"},
}

func TestSolution(t *testing.T) {
	t.Parallel()
	for i, tt := range cases {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.a, tt.in.b)
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

func Example_main2() {
	c, _ := ioutil.ReadFile("./input2.txt")
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
