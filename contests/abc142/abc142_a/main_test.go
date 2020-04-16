package main

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id    int
	input int
	want  float64
}{
	{id: 1, input: 4, want: 0.5000000000},
	{id: 2, input: 5, want: 0.6000000000},
	{id: 3, input: 1, want: 1.0000000000},
}

func TestOddPercent(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := oddPercent(tt.input)
			if got != tt.want {
				t.Errorf("%f, want: %f", got, tt.want)
			}
		})
	}
}
