package main

import (
	"fmt"
	"testing"
)

type input struct {
	n int
	a []int
}

var cases = []struct {
	id    int
	input input
	want  int
}{
	{id: 1, input: input{n: 3, a: []int{2, 1, 2}}, want: 1},
	{id: 2, input: input{n: 3, a: []int{2, 2, 2}}, want: -1},
	{id: 3, input: input{n: 10, a: []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}}, want: 7},
	{id: 4, input: input{n: 1, a: []int{1}}, want: 0},
}

func TestBrickBreak(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := brickBreak(tt.input.n, tt.input.a)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
