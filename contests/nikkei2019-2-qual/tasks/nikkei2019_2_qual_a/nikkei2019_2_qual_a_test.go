package main

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id    uint
	input uint
	want  uint
}{
	{id: 1, input: 4, want: 1},
	{id: 2, input: 999999, want: 499999},
	{id: 3, input: 5, want: 2},
	{id: 4, input: 6, want: 2},
}

func TestSumOfTwoIntegers(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintln(tt.id), func(t *testing.T) {
			got := sumOfTwoIntegers(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
