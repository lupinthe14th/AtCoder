package main

import (
	"fmt"
	"reflect"
	"testing"
)

var cases = []struct {
	id     int
	inputN int
	inputH []int
	want   []int
}{
	{id: 1, inputN: 3, inputH: []int{2, 3, 1}, want: []int{3, 1, 2}},
	{id: 2, inputN: 5, inputH: []int{1, 2, 3, 4, 5}, want: []int{1, 2, 3, 4, 5}},
	{id: 3, inputN: 8, inputH: []int{8, 2, 7, 3, 4, 5, 6, 1}, want: []int{8, 2, 4, 5, 6, 7, 3, 1}},
}

func TestGo2School(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := go2School(tt.inputN, tt.inputH)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
