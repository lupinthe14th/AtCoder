package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

type Case struct {
	input, want string
}

var err error
var cases = []Case{
	{input: "FisHDoGCaTAAAaAAbCAC", want: "AAAaAAbCACCaTDoGFisH"},
}

func TestSolution(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.input)
			if got != tt.want {
				t.Errorf("input: %s, got: %s, want: %s", tt.input, got, tt.want)
			}
		})
	}
}

func Example_main() {
	s := strings.NewReader("AAAAAjhfgaBCsahdfakGZsZGdEAA")
	b, _ := ioutil.ReadAll(s)
	inr, inw, _ := os.Pipe()
	orgStdin := os.Stdin
	inw.Write(b)
	inw.Close()
	os.Stdin = orgStdin
	os.Stdin = inr

	main()

	os.Stdin = orgStdin

	// Output: AAAAAAAjhfgaBCsahdfakGGdEZsZ
}
