package main

import (
	"fmt"
	"testing"
)

type input struct {
	s, t string
}

var cases = []struct {
	id    int
	input input
	want  string
}{
	{id: 1, input: input{s: "oder", t: "atc"}, want: "atcoder"},
	{id: 2, input: input{s: "humu", t: "humu"}, want: "humuhumu"},
}

func TestSolution(t *testing.T) {
	t.Parallel()
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := solution(tt.input.s, tt.input.t)
			if got != tt.want {
				t.Errorf("%s, want: %s", got, tt.want)
			}
		})
	}
}
