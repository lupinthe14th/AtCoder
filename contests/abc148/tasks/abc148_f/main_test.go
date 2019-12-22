package main

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id    int
	input int
	want  int
}{
	{id: 1, input: 12, want: 1},
	{id: 2, input: 5, want: 0},
	{id: 3, input: 1000000000000000000, want: 124999999999999995},
}

func TestDoubleFactorial(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := doubleFactorial(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
