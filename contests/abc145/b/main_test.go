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
	want  bool
}{
	{id: 1, input: input{n: 6, s: "abcabc"}, want: true},
	{id: 2, input: input{n: 6, s: "abcadc"}, want: false},
	{id: 3, input: input{n: 1, s: "z"}, want: false},
	{id: 4, input: input{n: 100, s: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}, want: true},
	{id: 5, input: input{n: 100, s: "vaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}, want: false},
}

func TestCircle(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := echo(tt.input.n, tt.input.s)
			if got != tt.want {
				t.Errorf("%t, want: %t", got, tt.want)
			}
		})
	}
}
