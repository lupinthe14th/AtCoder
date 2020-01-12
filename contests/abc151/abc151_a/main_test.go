package main

import (
	"fmt"
	"testing"
)

type Case struct {
	input string
	want  string
}

var cases = []Case{
	{input: "a", want: "b"},
	{input: "y", want: "z"},
}

func TestSolution(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.input)
			if got != tt.want {
				t.Errorf("%s, want: %s", got, tt.want)
			}
		})
	}
}
