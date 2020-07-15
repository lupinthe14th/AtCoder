package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type in struct {
	n    int
	nums []int
}

type Case struct {
	in   in
	want string
}

var cases = []Case{
	{in: in{n: 5, nums: []int{6, 7, 9, 10, 31}}, want: "APPROVED"},
}

func TestSolution(t *testing.T) {
	t.Parallel()
	for i, tt := range cases {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.n, tt.in.nums)
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

	// Output: DENIED
}
