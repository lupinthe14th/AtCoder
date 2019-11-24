package main

import (
	"fmt"
	"testing"
)

type input struct {
	n int
	s string
}

var cases = []struct {
	id    int
	input input
	want  string
}{
	{id: 1, input: input{n: 2, s: "ABCXYZ"}, want: "CDEZAB"},
	{id: 2, input: input{n: 0, s: "ABCXYZ"}, want: "ABCXYZ"},
	{id: 3, input: input{n: 13, s: "ABCDEFGHIJKLMNOPQRSTUVWXYZ"}, want: "NOPQRSTUVWXYZABCDEFGHIJKLM"},
}

func TestRotN(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := rotN(tt.input.n, tt.input.s)
			if got != tt.want {
				t.Errorf("%s, want: %s", got, tt.want)
			}
		})
	}
}
