package main

import (
	"fmt"
	"testing"
)

type input struct {
	one   int
	two   int
	three int
}

var cases = []struct {
	id    int
	input input
	want  string
}{
	{id: 1, input: input{one: 5, two: 7, three: 9}, want: "win"},
	{id: 2, input: input{one: 13, two: 7, three: 2}, want: "bust"},
}

func TestBlackjack(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := blackjack(tt.input.one, tt.input.two, tt.input.three)
			if got != tt.want {
				t.Errorf("%s, want: %s", got, tt.want)
			}
		})
	}
}
