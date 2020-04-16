package main

import (
	"fmt"
	"testing"
)

type input struct {
	n int
	s string
	t string
}

var cases = []struct {
	id    int
	input input
	want  string
}{
	{id: 1, input: input{n: 2, s: "ip", t: "cc"}, want: "icpc"},
	{id: 2, input: input{n: 8, s: "hmhmnknk", t: "uuuuuuuu"}, want: "humuhumunukunuku"},
	{id: 3, input: input{n: 5, s: "aaaaa", t: "aaaaa"}, want: "aaaaaaaaaa"},
}

func TestStringsWithTheSameLength(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := stringsWithTheSameLength(tt.input.n, tt.input.s, tt.input.t)
			if got != tt.want {
				t.Errorf("%s, want: %s", got, tt.want)
			}
		})
	}
}
