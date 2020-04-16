package main

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id    int
	input int
	want  bool
}{
	{id: 1, input: 615, want: true},
	{id: 2, input: 217, want: false},
}

func TestOneOOToOneOFive(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := oneOOToOneOFive(tt.input)
			if got != tt.want {
				t.Errorf("%t, want: %t", got, tt.want)
			}
		})
	}
}
