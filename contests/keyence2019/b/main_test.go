package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in, want string
	}{
		{in: "keyofscience", want: "YES"},
		{in: "mpyszsbznf", want: "NO"},
		{in: "keyence", want: "YES"},
		{in: "keyofenscice", want: "NO"},
		{in: "keyencbbzsqzezyvefbhtwriljcvrymirjmbljehlslmrxslnspsiujxiznmlyfcpe", want: "YES"},
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
	fd, _ := os.Open(filepath.Join("testdata", "input.txt"))
	orgStdin := os.Stdin
	os.Stdin = fd
	defer func() {
		os.Stdin = orgStdin
		fd.Close()
	}()

	main()

	// Output: NO
}
