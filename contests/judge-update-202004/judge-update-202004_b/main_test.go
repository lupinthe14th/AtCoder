package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

type in struct {
	n    int
	r, b []int
}

type Case struct {
	in   in
	want []int
}

var cases = []Case{
	{in: in{n: 4, r: []int{6, 2}, b: []int{10, 4}}, want: []int{2, 6, 4, 10}},
}

func TestSolution(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in.n, tt.in.r, tt.in.b)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %+v, got: %d, want: %d", tt.in, got, tt.want)
			}
		})
	}
}

func Example_main() {
	c, _ := ioutil.ReadFile("./input.txt")
	inr, inw, _ := os.Pipe()
	orgStdin := os.Stdin
	inw.Write(c)
	inw.Close()
	os.Stdin = orgStdin
	defer func() { os.Stdin = orgStdin }()
	os.Stdin = inr

	main()

	// Output:
	// 5
	// 7
}
