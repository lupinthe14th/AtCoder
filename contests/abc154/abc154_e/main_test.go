package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type input struct {
	n string
	k int
}

type Case struct {
	input input
	want  int
}

var cases = []Case{
	{input: input{n: "100", k: 1}, want: 19},
	{input: input{n: "25", k: 2}, want: 14},
	{input: input{n: "314159", k: 2}, want: 937},
	{input: input{n: "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", k: 3}, want: 117879300},
}

func TestSolution(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.input.n, tt.input.k)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
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
	os.Stdin = inr

	main()

	os.Stdin = orgStdin

	// Output: 117879300
}
