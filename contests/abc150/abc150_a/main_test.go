package main

import (
	"fmt"
	"testing"
)

type input struct {
	k, x int
}

type Case struct {
	input input
	want  bool
}

var cases = []Case{
	{input: input{k: 2, x: 900}, want: true},
	{input: input{k: 1, x: 501}, want: false},
}

func TestSolution(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.input.k, tt.input.x)
			if got != tt.want {
				t.Errorf("%t, want: %t", got, tt.want)
			}
		})
	}
}
