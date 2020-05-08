package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type in struct {
		n, a, b int
		s       string
	}

	var cases = []struct {
		in   in
		want []string
	}{
		{in: in{n: 10, a: 2, b: 3, s: "abccabaabb"}, want: []string{"Yes", "Yes", "No", "No", "Yes", "Yes", "Yes", "No", "No", "No"}},
		{in: in{n: 5, a: 2, b: 2, s: "ccccc"}, want: []string{"No", "No", "No", "No", "No"}},
	}
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solution(tt.in.n, tt.in.a, tt.in.b, tt.in.s)
			if !reflect.DeepEqual(got, tt.want) {
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
	inw.Close()
	os.Stdin = orgStdin
	defer func() { os.Stdin = orgStdin }()
	os.Stdin = inr

	main()

	// Output:
	// No
	// Yes
	// Yes
	// Yes
	// Yes
	// No
	// Yes
	// Yes
	// No
	// Yes
	// No
	// No
}
