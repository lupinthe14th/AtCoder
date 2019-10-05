package main

import (
	"fmt"
	"testing"
)

type input struct {
	s string
	k int
}

var cases = []struct {
	id    int
	input input
	want  int
}{
	{id: 1, input: input{s: "issii", k: 2}, want: 4},
	{id: 2, input: input{s: "qq", k: 81}, want: 81},
	{id: 3, input: input{s: "cooooooooonteeeeeeeeeest", k: 999993333}, want: 8999939997},
	{id: 4, input: input{s: "i", k: 2}, want: 1},
}

func TestConnectionAndDisconnection(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintln(tt.id), func(t *testing.T) {
			got := connectionAndDisconnection(tt.input.s, tt.input.k)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
