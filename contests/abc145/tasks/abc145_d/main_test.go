package main

import (
	"fmt"
	"testing"
)

type input struct {
	x, y int
}

var cases = []struct {
	id    int
	input input
	want  int
}{
	{id: 1, input: input{x: 3, y: 3}, want: 2},
	{id: 2, input: input{x: 2, y: 2}, want: 0},
	{id: 3, input: input{x: 999999, y: 999999}, want: 151840682},
}

func TestKnight(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := knight(tt.input.x, tt.input.y)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
