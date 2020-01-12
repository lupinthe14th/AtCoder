package main

import (
	"fmt"
	"reflect"
	"testing"
)

type input struct {
	n, m int
	r    []result
}

type Case struct {
	input input
	want  ans
}

var cases = []Case{
	{input: input{n: 2, m: 5, r: []result{{p: 1, s: "WA"}, {p: 1, s: "AC"}, {p: 2, s: "WA"}, {p: 2, s: "AC"}, {p: 2, s: "WA"}}}, want: ans{c: 2, p: 2}},
	{input: input{n: 100000, m: 3, r: []result{{p: 7777, s: "AC"}, {p: 7777, s: "AC"}, {p: 7777, s: "AC"}}}, want: ans{c: 1, p: 0}},
	{input: input{n: 6, m: 0, r: []result{}}, want: ans{c: 0, p: 0}},
}

func TestSolution(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.input.n, tt.input.m, tt.input.r)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}
