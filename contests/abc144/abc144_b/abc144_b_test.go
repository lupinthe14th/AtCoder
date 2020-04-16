package main

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id    int
	input int
	want  string
}{
	{id: 1, input: 10, want: "Yes"},
	{id: 2, input: 50, want: "No"},
	{id: 3, input: 81, want: "Yes"},
}

func TestIsMultiplicationTable(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintln(tt.id), func(t *testing.T) {
			actual := isMultiplicationTable(tt.input)
			if actual != tt.want {
				t.Errorf("%s, want: %s", actual, tt.want)
			}
		})
	}
}
