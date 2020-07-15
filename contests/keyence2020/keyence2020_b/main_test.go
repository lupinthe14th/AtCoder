package main

import (
	"fmt"
	"testing"
)

type input struct {
	n      int
	matrix [][]int
}

type Case struct {
	input input
	want  int
}

var cases = []Case{
	{input: input{n: 4, matrix: [][]int{{2, 4}, {4, 3}, {9, 3}, {100, 5}}}, want: 3},
	{input: input{n: 2, matrix: [][]int{{8, 20}, {1, 10}}}, want: 1},
	{input: input{n: 5, matrix: [][]int{{10, 1}, {2, 1}, {4, 1}, {6, 1}, {8, 1}}}, want: 5},
}

func TestSolution(t *testing.T) {
	t.Parallel()
	for i, tt := range cases {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.input.n, tt.input.matrix)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
