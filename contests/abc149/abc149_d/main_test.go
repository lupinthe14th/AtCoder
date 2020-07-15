package main

import (
	"fmt"
	"testing"
)

type input struct {
	n, k    int
	r, s, p int
	t       string
}

type Case struct {
	input input
	want  int
}

var cases = []Case{
	{input: input{n: 5, k: 2, r: 8, s: 7, p: 6, t: "rsrpr"}, want: 27},
	{input: input{n: 7, k: 1, r: 100, s: 10, p: 1, t: "ssssppr"}, want: 211},
	{input: input{n: 30, k: 5, r: 325, s: 234, p: 123, t: "rspsspspsrpspsppprpsprpssprpsr"}, want: 4996},
}

func TestSolution(t *testing.T) {
	t.Parallel()
	for i, tt := range cases {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.input.n, tt.input.k, tt.input.r, tt.input.s, tt.input.p, tt.input.t)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
