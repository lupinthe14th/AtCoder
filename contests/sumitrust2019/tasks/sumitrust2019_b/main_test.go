package main

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id      int
	input   int
	want    int
	wanterr string
}{
	{id: 1, input: 432, want: 400},
	{id: 2, input: 1079, wanterr: ":("},
	{id: 3, input: 1001, want: 927},
}

func TestTaxRate(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got, err := taxRate(tt.input)
			if got == 0 && err != tt.wanterr {
				t.Errorf("%s, want: %s", err, tt.wanterr)
			}
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
