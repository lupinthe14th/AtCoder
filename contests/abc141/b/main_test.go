package main

import (
	"fmt"
	"os"
	"testing"
)

func TestIsTapDanceEasy(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{in: "U", want: "Yes"},
		{in: "RUDLUDR", want: "Yes"},
		{in: "DULL", want: "No"},
		{in: "UUUUUUUUUUUUUUU", want: "Yes"},
		{in: "UUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUU", want: "Yes"},
		{in: "ULURU", want: "No"},
		{in: "RDULULDURURLRDULRLR", want: "Yes"},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
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

	// Output: Yes
}
