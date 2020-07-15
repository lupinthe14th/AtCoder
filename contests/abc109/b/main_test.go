package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Parallel()
	type in struct {
		n int
		w []string
	}
	tests := []struct {
		in   in
		want string
	}{
		{in: in{n: 4, w: []string{"hoge", "english", "hoge", "enigma"}}, want: "No"},
		{in: in{n: 8, w: []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaa", "aaaaaaa"}}, want: "No"},
		{in: in{n: 3, w: []string{"abc", "arc", "agc"}}, want: "No"},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.n, tt.in.w)
			if got != tt.want {
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

	// Output: Yes
}
