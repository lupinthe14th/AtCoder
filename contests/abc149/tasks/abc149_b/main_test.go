package main

import (
	"fmt"
	"reflect"
	"testing"
)

type input struct {
	a, b, k uint64
}

var cases = []struct {
	id    int
	input input
	want  []uint64
}{
	{id: 1, input: input{a: 2, b: 3, k: 3}, want: []uint64{0, 2}},
	{id: 2, input: input{a: 500000000000, b: 500000000000, k: 1000000000000}, want: []uint64{0, 0}},
}

func TestSolution(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := solution(tt.input.a, tt.input.b, tt.input.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}
