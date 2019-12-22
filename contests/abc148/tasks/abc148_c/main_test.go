package main

import (
	"fmt"
	"testing"
)

type input struct {
	a, b int
}

var cases = []struct {
	id    int
	input input
	want  int
}{
	{id: 1, input: input{a: 2, b: 3}, want: 6},
	{id: 2, input: input{a: 123, b: 456}, want: 18696},
	{id: 3, input: input{a: 100000, b: 99999}, want: 9999900000},
}

func TestSnack(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := snack(tt.input.a, tt.input.b)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
