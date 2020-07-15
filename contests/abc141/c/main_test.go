package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Parallel()
	type in struct {
		n, k, q int
		a       []int
	}

	var cases = []struct {
		in   in
		want []string
	}{
		{in: in{n: 6, k: 3, q: 4, a: []int{3, 1, 3, 2}}, want: []string{"No", "No", "Yes", "No", "No", "No"}},
		{in: in{n: 6, k: 5, q: 4, a: []int{3, 1, 3, 2}}, want: []string{"Yes", "Yes", "Yes", "Yes", "Yes", "Yes"}},
	}
	for i, tt := range cases {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.n, tt.in.k, tt.in.q, tt.in.a)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("in: %+v, got: %v, want: %v", tt.in, got, tt.want)
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

	// Output:
	// No
	// No
	// No
	// No
	// Yes
	// No
	// No
	// No
	// Yes
	// No
}
