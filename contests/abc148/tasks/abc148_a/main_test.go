package main

import (
	"fmt"
	"testing"
)

type input struct {
	a int
	b int
}

var cases = []struct {
	id    int
	input input
	want  int
}{
	{id: 1, input: input{a: 3, b: 1}, want: 2},
	{id: 2, input: input{a: 1, b: 2}, want: 3},
	{id: 3, input: input{a: 2, b: 3}, want: 1},
}

func TestRoundOne(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := roundOne(tt.input.a, tt.input.b)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
