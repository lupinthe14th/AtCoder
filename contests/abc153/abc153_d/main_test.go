package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

type Case struct {
	input, want int
}

var cases = []Case{
	{input: 2, want: 3},
	{input: 4, want: 7},
	{input: 8, want: 15},
	{input: 16, want: 31},
	{input: 32, want: 63},
	{input: 100, want: 127},
	{input: 160, want: 255},
}

func TestSolution(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}

func Example_main() {
	s := strings.NewReader("1000000000000")
	b, _ := ioutil.ReadAll(s)
	inr, inw, _ := os.Pipe()
	orgStdin := os.Stdin
	inw.Write(b)
	inw.Close()
	os.Stdin = orgStdin
	os.Stdin = inr

	main()

	os.Stdin = orgStdin

	// Output: 1099511627775
}
