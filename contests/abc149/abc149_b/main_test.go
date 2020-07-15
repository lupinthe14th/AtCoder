package main

import (
	"fmt"
	"reflect"
	"testing"
)

type input struct {
	a, b, k int64
}

var cases = []struct {
	id    int
	input input
	want  []int64
}{
	{id: 1, input: input{a: 2, b: 3, k: 3}, want: []int64{0, 2}},
	{id: 2, input: input{a: 500000000000, b: 500000000000, k: 1000000000000}, want: []int64{0, 0}},
	{id: 3, input: input{a: 0, b: 0, k: 0}, want: []int64{0, 0}},
	{id: 4, input: input{a: 2, b: 3, k: 2}, want: []int64{0, 3}},
	{id: 5, input: input{a: 5, b: 3, k: 3}, want: []int64{2, 3}},
	{id: 6, input: input{a: 5, b: 3, k: 9}, want: []int64{0, 0}},
}

func TestSolution(t *testing.T) {
	t.Parallel()
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := solution(tt.input.a, tt.input.b, tt.input.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}
