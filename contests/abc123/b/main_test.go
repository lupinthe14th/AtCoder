package main

import (
	"fmt"
	"os"
	"testing"
)

func TestSolution(t *testing.T) {
	var tests = []struct {
		in   Order
		want int
	}{
		{in: Order{29, 20, 7, 35, 120}, want: 215},
		{in: Order{123, 123, 123, 123, 123}, want: 643},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}

func Example_main() {
	fd, _ := os.Open("./input.txt")
	orgStdin := os.Stdin
	os.Stdin = orgStdin
	os.Stdin = fd
	defer func() {
		os.Stdin = orgStdin
		fd.Close()
	}()

	main()

	// Output: 481
}
