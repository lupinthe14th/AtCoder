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
		n, k int
	}
	var cases = []struct {
		in   in
		want int
	}{
		{in: in{n: 3, k: 2}, want: 10},
		{in: in{n: 141421, k: 35623}, want: 220280457},
	}
	for i, tt := range cases {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.n, tt.in.k)
			if got != tt.want {
				t.Errorf("in: %+v, got: %v, want: %v", tt.in, got, tt.want)
			}
		})
	}
}

func Example_main() {
	c, _ := ioutil.ReadFile("./input.txt")
	inr, inw, _ := os.Pipe()
	orgStdin := os.Stdin
	inw.Write(c)
	os.Stdin = orgStdin
	defer func() {
		inw.Close()
		os.Stdin = orgStdin
	}()
	os.Stdin = inr

	main()

	// Output: 1
}
