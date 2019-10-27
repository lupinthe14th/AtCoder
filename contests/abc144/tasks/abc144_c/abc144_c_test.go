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
	{id: 1, input: 10, want: 5},
	{id: 2, input: 50, want: 13},
	{id: 3, input: 10000000019, want: 10000000018},
}

func TestWalkOnMultiplicationTable(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintln(tt.id), func(t *testing.T) {
			got := walkOnMultiplicationTable(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
