package main

import (
	"fmt"
	"reflect"
	"testing"
)

type input struct {
	n, k, s int
}

type Case struct {
	input input
	want  []int
}

var cases = []Case{
	{input: input{n: 4, k: 2, s: 3}, want: []int{3, 3, 4, 4}},
	{input: input{n: 5, k: 3, s: 100}, want: []int{100, 100, 100, 101, 101}},
}

func TestSolution(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.input.n, tt.input.k, tt.input.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}
