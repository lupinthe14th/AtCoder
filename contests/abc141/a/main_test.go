package main

import (
	"fmt"
	"os"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{in: "Sunny", want: "Cloudy"},
		{in: "Rainy", want: "Sunny"},
		{in: "Cloudy", want: "Rainy"},
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
	os.Stdin = fd
	defer func() {
		os.Stdin = orgStdin
		fd.Close()
	}()

	main()

	// Output: Sunny
}
