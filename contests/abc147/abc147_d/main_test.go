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
	{id: 1, input: input{n: 3, a: []int{1, 2, 3}}, want: 6},
	{id: 2, input: input{n: 10, a: []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}}, want: 237},
	{id: 3, input: input{n: 10, a: []int{3, 14, 159, 2653, 58979, 323846, 2643383, 27950288, 419716939, 9375105820}}, want: 103715602},
}

func TestXorSum4(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := xorSum4(tt.input.n, tt.input.a)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
