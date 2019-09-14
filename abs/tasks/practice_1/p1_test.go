package main

import (
	"testing"
)

var cases = []struct {
	a, b, c int
	s, want string
}{
	{a: 1, b: 2, c: 3, s: "test", want: "6 test"},
}

func TestPrintSumAndWord(t *testing.T) {
	for _, tt := range cases {
		actual := printSumAndWord(tt.a, tt.b, tt.c, tt.s)
		if actual != tt.want {
			t.Errorf("%s, want: %s", actual, tt.want)
		}
	}
}
