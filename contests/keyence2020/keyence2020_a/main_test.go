package main

import (
	"fmt"
	"testing"
)

type input struct {
	h, w, n int
}

type Case struct {
	input input
	want  int
}

var cases = []Case{
	{input: input{h: 3, w: 7, n: 10}, want: 2},
	{input: input{h: 14, w: 12, n: 112}, want: 8},
	{input: input{h: 2, w: 100, n: 200}, want: 2},
}

func TestSolution(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.input.h, tt.input.w, tt.input.n)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
