package main

import (
	"fmt"
	"reflect"
	"testing"
)

type input struct {
	h, w, a, b int
}

var cases = []struct {
	id    int
	input input
	want  [][]int
}{
	{id: 1, input: input{h: 3, w: 3, a: 1, b: 1}, want: [][]int{[]int{1, 0, 0}, []int{0, 1, 0}, []int{0, 0, 1}}},
}

func TestMatrix(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := Matrix(tt.input.h, tt.input.w, tt.input.a, tt.input.b)
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}
