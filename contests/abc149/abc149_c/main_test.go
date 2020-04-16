package main

import (
	"fmt"
	"testing"
)

type Case struct {
	input int
	want  int
}

var cases = []Case{
	{input: 20, want: 23},
	{input: 2, want: 2},
	{input: 99992, want: 100003},
	{input: 1, want: 2},
}

func TestSolution(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := trialDivision(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
