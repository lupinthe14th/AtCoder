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
	{id: 1, input: "redcoder", want: 1},
	{id: 2, input: "vvvvvv", want: 0},
	{id: 3, input: "abcdabc", want: 2},
}

func TestPalindromePhilia(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := palindromePhilia(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
