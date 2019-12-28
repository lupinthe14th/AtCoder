package main

import (
	"fmt"
	"testing"
)

type input struct {
	n, a, b uint64
}

var cases = []struct {
	id    int
	input input
	want  uint64
}{
	{id: 1, input: input{n: 5, a: 2, b: 4}, want: 1},
	{id: 2, input: input{n: 5, a: 2, b: 3}, want: 2},
	{id: 3, input: input{n: 5, a: 1, b: 3}, want: 1},
	{id: 4, input: input{n: 6, a: 1, b: 3}, want: 1},
	{id: 5, input: input{n: 6, a: 2, b: 4}, want: 1},
	{id: 6, input: input{n: 6, a: 2, b: 3}, want: 2},
	{id: 7, input: input{n: 10, a: 2, b: 3}, want: 2},
}

func TestTableTennisTraining(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := tableTennisTraining(tt.input.n, tt.input.a, tt.input.b)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
