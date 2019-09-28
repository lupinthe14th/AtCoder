package main

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id     int
	inputN int
	inputK int
	inputH []int
	want   int
}{
	{id: 1, inputN: 4, inputK: 150, inputH: []int{150, 140, 100, 200}, want: 2},
	{id: 2, inputN: 1, inputK: 500, inputH: []int{499}, want: 0},
	{id: 3, inputN: 5, inputK: 1, inputH: []int{100, 200, 300, 400, 500}, want: 5},
}

func TestRollerCoaster(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := rollerCoaster(tt.inputN, tt.inputK, tt.inputH)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
