package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestSolution(t *testing.T) {
	var cases = []struct {
		in   int
		want float64
	}{
		{in: 1, want: 6.28318530717958623200},
	}
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in)
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

	// Output: 458.6725274241098
}
