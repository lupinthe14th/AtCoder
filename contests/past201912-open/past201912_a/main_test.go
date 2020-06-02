package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

type want struct {
	w int
	e error
}

type Case struct {
	input string
	want  want
}

var err error
var cases = []Case{
	{input: "678", want: want{w: 1356}},
	{input: "abc", want: want{e: err}},
	{input: "0x8", want: want{e: err}},
	{input: "012", want: want{w: 24}},
}

func TestSolution(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got, e := solution(tt.input)
			if tt.want.e != nil && e == nil {
				t.Errorf("%s, want.e: %s", e, tt.want.e)
			} else if got != tt.want.w {
				t.Errorf("%d, want.w: %d", got, tt.want.w)
			}
		})
	}
}

func Example_main() {
	s := strings.NewReader("abs")
	b, _ := ioutil.ReadAll(s)
	inr, inw, _ := os.Pipe()
	orgStdin := os.Stdin
	inw.Write(b)
	inw.Close()
	os.Stdin = orgStdin
	os.Stdin = inr

	main()

	os.Stdin = orgStdin

	// Output: error
}
