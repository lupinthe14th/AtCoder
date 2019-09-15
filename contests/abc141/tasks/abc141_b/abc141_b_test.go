package main

import (
	"fmt"
	"testing"
)

var cases = []struct {
	input string
	want  bool
}{
	{input: "U", want: true},
	{input: "RUDLUDR", want: true},
	{input: "DULL", want: false},
	{input: "UUUUUUUUUUUUUUU", want: true},
	{input: "UUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUU", want: true},
	{input: "ULURU", want: false},
	{input: "RDULULDURURLRDULRLR", want: true},
}

func TestIsTapDanceEasy(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintln(tt.input), func(t *testing.T) {
			actual := isTapDanceEasy(tt.input)
			if actual != tt.want {
				t.Errorf("%t, want: %t", actual, tt.want)
			}
		})
	}
}
