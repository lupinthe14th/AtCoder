package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Parallel()
	type in struct {
		x1, y1, x2, y2 int
	}
	type want struct {
		x3, y3, x4, y4 int
	}
	var cases = []struct {
		in   in
		want want
	}{
		{in: in{x1: 0, y1: 0, x2: 0, y2: 1}, want: want{x3: -1, y3: 1, x4: -1, y4: 0}},
		{in: in{x1: 2, y1: 3, x2: 6, y2: 6}, want: want{x3: 3, y3: 10, x4: -1, y4: 7}},
	}
	for i, tt := range cases {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			x3, y3, x4, y4 := solution(tt.in.x1, tt.in.y1, tt.in.x2, tt.in.y2)
			if !(x3 == tt.want.x3 && y3 == tt.want.y3 && x4 == tt.want.x4 && y4 == tt.want.y4) {
				t.Errorf("in: %+v, x3: %v, y3: %v, x4: %v, y4: %v, want: %v", tt.in, x3, y3, x4, y4, tt.want)
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

	// Output: -126 -64 -36 -131
}
