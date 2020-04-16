package main

import (
	"fmt"
	"testing"
)

type input struct {
	m1, d1, m2, d2 int
}

var cases = []struct {
	id    int
	input input
	want  int
}{
	{id: 1, input: input{m1: 11, d1: 16, m2: 11, d2: 17}, want: 0},
	{id: 2, input: input{m1: 11, d1: 30, m2: 12, d2: 1}, want: 1},
}

func TestLastDayOfTheMonth(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := lastDayOfTheMonth(tt.input.m1, tt.input.d1, tt.input.m2, tt.input.d2)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
