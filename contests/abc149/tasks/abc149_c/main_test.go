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
	{id: 1, input: 20, want: 23},
	{id: 2, input: 2, want: 2},
	{id: 3, input: 99992, want: 100003},
}

func TestSolution(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := trialDivision(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
