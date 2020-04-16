package main

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id    int
	input string
	want  int
}{
	{id: 1, input: "SAT", want: 1},
	{id: 2, input: "SUN", want: 7},
	{id: 3, input: "MON", want: 6},
	{id: 4, input: "TUE", want: 5},
	{id: 5, input: "WED", want: 4},
	{id: 6, input: "THU", want: 3},
	{id: 7, input: "FRI", want: 2},
}

func TestCircle(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := canNotWaitForHoliday(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
