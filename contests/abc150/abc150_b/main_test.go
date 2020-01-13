package main

import (
	"fmt"
	"testing"
)

type input struct {
	n int
	s string
}

type Case struct {
	input input
	want  int
}

var cases = []Case{
	{input: input{n: 10, s: "ZABCDBABCQ"}, want: 2},
}

func TestSolution(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.input.n, tt.input.s)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
