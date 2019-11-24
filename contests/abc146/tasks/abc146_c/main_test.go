package main

import (
	"fmt"
	"testing"
)

type input struct {
	a, b, x int
}

var cases = []struct {
	id    int
	input input
	want  int
}{
	{id: 1, input: input{a: 10, b: 7, x: 100}, want: 9},
	{id: 2, input: input{a: 2, b: 1, x: 100000000000}, want: 1000000000},
	{id: 3, input: input{a: 1000000000, b: 1000000000, x: 100}, want: 0},
	{id: 4, input: input{a: 1234, b: 56789, x: 314159265}, want: 254309},
}

func TestBuyAnInteger(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := buyAnInteger(tt.input.a, tt.input.b, tt.input.x)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
