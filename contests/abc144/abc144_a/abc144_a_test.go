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
	{id: 1, input: input{a: 2, b: 5}, want: 10},
	{id: 2, input: input{a: 2, b: 10}, want: -1},
	{id: 3, input: input{a: 9, b: 9}, want: 81},
}

func TestMultiplicationTable(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintln(tt.id), func(t *testing.T) {
			actual := multiplicationTable(tt.input.a, tt.input.b)
			if actual != tt.want {
				t.Errorf("%d, want: %d", actual, tt.want)
			}
		})
	}
}
