package main

import (
	"fmt"
	"testing"
)

type input struct {
	n int
	a []int
	b []int
}

var cases = []struct {
	id    int
	input input
	want  bool
}{
	{id: 1, input: input{n: 3, a: []int{1, 3, 2}, b: []int{1, 2, 3}}, want: true},
	{id: 2, input: input{n: 3, a: []int{1, 2, 3}, b: []int{2, 2, 2}}, want: false},
	{id: 3, input: input{n: 6, a: []int{3, 1, 2, 6, 3, 4}, b: []int{2, 2, 8, 3, 4, 3}}, want: true},
	{id: 4, input: input{n: 5, a: []int{3, 1, 2, 5, 4}, b: []int{3, 2, 4, 1, 5}}, want: true},
}

func TestSwaps(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintln(tt.id), func(t *testing.T) {
			got := swaps(tt.input.n, tt.input.a, tt.input.b)
			if got != tt.want {
				t.Errorf("%t, want: %t", got, tt.want)
			}
		})
	}
}
