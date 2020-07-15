package main

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Parallel()
	type in struct {
		n           int
		restaurants []restaurant
	}
	tests := []struct {
		in   in
		want []int
	}{
		{in: in{n: 6, restaurants: []restaurant{{id: 1, city: "khabarovsk", rate: 20}, {id: 2, city: "moscow", rate: 10}, {id: 3, city: "kazan", rate: 50}, {id: 4, city: "kazan", rate: 35}, {id: 5, city: "moscow", rate: 60}, {id: 6, city: "khabarovsk", rate: 40}}}, want: []int{3, 4, 6, 1, 5, 2}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.n, tt.in.restaurants)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}

func Example_main() {
	fd, _ := os.Open(filepath.Join("testdata", "input.txt"))
	orgStdin := os.Stdin
	os.Stdin = fd
	defer func() {
		os.Stdin = orgStdin
		fd.Close()
	}()

	main()

	// Output:
	// 10
	// 9
	// 8
	// 7
	// 6
	// 5
	// 4
	// 3
	// 2
	// 1
}
